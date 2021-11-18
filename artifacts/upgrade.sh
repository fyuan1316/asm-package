

# 设置集群名称
cluster=region
# 查看待升级的 asm 资源
kubectl get -n cpaas-system moduleinfo -l cpaas.io/cluster-name="${cluster}",cpaas.io/module-name=asm
# 查看待升级的 istio 资源
kubectl get -n cpaas-system moduleinfo -l cpaas.io/cluster-name="${cluster}",cpaas.io/module-name=istio


asm_mod="$(kubectl get -n cpaas-system moduleinfo -l cpaas.io/cluster-name="${cluster}",cpaas.io/module-name=asm -o=name)"

########################################################
# 平台升级，严格按照顺序执行，asm升级完毕后，才能开始istio升级！
########################################################
function isComplete() {
    local mod_name=$1
    local upgrade_version=$2
    spec_version="$(kubectl get -n cpaas-system "${mod_name}" -ojsonpath='{.spec.version}')"
    status_version="$(kubectl get -n cpaas-system "${mod_name}" -ojsonpath='{.status.version}')"
    phase="$(kubectl get -n cpaas-system "${mod_name}" -ojsonpath='{.status.phase}')"
    echo "${spec_version} ${status_version} ${phase}"
    if [[ "${phase}" == "Upgrading" ]]; then
        echo "升级中"
    elif [[ "${spec_version}" == "${status_version}" ]] &&
        [[ "${phase}" == "Running" ]]; then

        if [[ "${upgrade_version}" == "${spec_version}" ]]; then
            echo "升级完毕"
        else
            echo "未升级到预期版本，预期为${upgrade_version}"
        fi
    fi
}

# 根据现场实际情况设置
# 集群名称
cluster=region
##########
# 第一步
##########
asm_mod="$(kubectl get -n cpaas-system moduleinfo -l cpaas.io/cluster-name="${cluster}",cpaas.io/module-name=asm -o=name)"
# asm升级的目标版本
asm_version="$(kubectl get packagemanifest asm-operator -o=jsonpath='{.status.channels[0].currentCSV}' | cut -d . -f 2-)"

# 执行升级asm
kubectl patch -n cpaas-system "${asm_mod}" --type='json' -p="[{'op': 'replace', 'path': '/spec/version', 'value': ${asm_version}]"

# 检查asm升级是否完毕
isComplete "${asm_mod}" "${asm_version}"

##########
# 第二步
##########
istio_mod="$(kubectl get -n cpaas-system moduleinfo -l cpaas.io/cluster-name="${cluster}",cpaas.io/module-name=istio -o=name)"
# istio升级的目标版本
istio_version=1.10.5

# 升级istio
kubectl patch -n cpaas-system "${istio_mod}" --type='json' -p="[{'op': 'replace', 'path': '/spec/version', 'value': ${istio_version}]"

# 检查istio升级是否完毕
isComplete "${istio_mod}" "${istio_version}"