
function check_has_upgrade_plan() {
    count="$(kubectl get -n istio-system installplan | grep false -c)"
    if [[ ${count} -gt 0 ]]; then
        echo "存在待升级的installplan, 当前OLM升级检测到csv版本:"
        arr="$(kubectl get -n istio-system installplan | grep false | awk '{print $1}')"
        for plan in ${arr[@]}; do
            v="$(kubectl get -n istio-system installplan $plan -ojsonpath='{.spec.clusterServiceVersionNames}')"
            echo "$v"
        done
        echo "请执行[check_csv]对照检查预期版本是否全部存在于当前检测到的版本."
        echo ""
    else
        echo "没找到可升级的installplan"
    fi
}

function check_csv() {
    asm_csv="$(kubectl get -n cpaas-system artifactversion "asm-operator.$asm_bundle_version" -ojsonpath='{.status.version}')"
    flagger_csv="$(kubectl get -n cpaas-system artifactversion "flagger-operator.$flagger_bundle_version" -ojsonpath='{.status.version}')"
    echo "升级所需asm_csv    : ${asm_csv}"
    echo "升级所需flagger_csv: ${flagger_csv}"
    echo ""
}

# 设置 asm和flagger的bundle版本
asm_bundle_version=v3.7-13-ge53b7de
flagger_bundle_version=v3.7-3-ga0a14d5

# 检查是否存在待升级的installplan
check_has_upgrade_plan
# 检查升级所需的Csv版本
check_csv