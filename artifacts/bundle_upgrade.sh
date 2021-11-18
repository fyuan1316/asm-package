#!/usr/bin/env bash
set -e

Asm_Bundle=asm-operator-bundle
Flagger_Bundle=flagger-operator-bundle
Asm_Bundle_Version=
Flagger_Bundle_Version=
localRegistry=

# sh bundle_upgrade.sh -a v3.7-13-ge53b7de -f v3.7-3-ga0a14d5 -r 192.168.131.248:60080
function usage() {
    echo "Usage: $0 -a <asm-bundle-version> -f <flagger-bundle-version> -r <registry>"
    echo "  -a <asm-bundle-version>    : The version of the ASM bundle to install"
    echo "  -f <flagger-bundle-version>: The version of the Flagger bundle to install"
    echo "  -r <registry>              : Registry"
    echo "  -h                        : Prints this help message"
    exit 1
}

check_bundle_version() {
    if [[ -z "$Asm_Bundle_Version" ]]; then
        echo "Please specify the ASM bundle version"
        usage
    fi
    if [[ -z "$Flagger_Bundle_Version" ]]; then
        echo "Please specify the Flagger bundle version"
        usage
    fi
    if [[ -z "$localRegistry" ]]; then
        echo "Please specify Local Registry"
        usage
    fi
}

function load() {
    if docker load -i "asm-bundle.tar"; then
        echo "asm-bundle.tar loaded successfully"
    else
        echo "asm-bundle.tar load failed"
        exit 1
    fi
}

# load and push bundle
function push() {
    local bundle_name=$1
    local bundle_version=$2
    local bundle_registry=$3

    img="$(docker images | grep "${bundle_name}" | grep "build-harbor.alauda.cn" | awk '{print $1":"$2}')"
    echo $img
    docker tag "${img}" "${bundle_registry}/asm/${bundle_name}:${bundle_version}" && docker push "${bundle_registry}/asm/${bundle_name}:${bundle_version}"
    echo "bundle $bundle_name pushed."

}
function sync_artifact() {
    local artifact_name=$1
    name=$(kubectl get artifact -A | grep ${artifact_name} | awk '{print $2}')
    curl -k -s -X PATCH -H "Accept: application/json, */*" \
        -H "Content-Type: application/merge-patch+json" \
        127.0.0.1:8001/apis/app.alauda.io/v1alpha1/namespaces/cpaas-system/artifacts/"${name}"/status \
        --data '{"status":{"synced":false}}'
}

function sync() {
    sync_artifact "asm-operator"
    sync_artifact "flagger-operator"
    echo ""
    kubectl delete pods -n cpaas-system -l app=catalog-operator
}
function proxyOn() {
    exist="$(ps -ef | grep "kubectl proxy" | grep -v color | awk '{print $2}')"
    if [[ -z "$exist" ]]; then
        echo "$exist"
        kubectl proxy &
    fi
    until [[ "$(curl -I -s -w %{http_code} 127.0.0.1:8001/version)" != "000" ]]; do
        echo "Waiting for kubectl proxy ready..."
        sleep 1
    done
}

while getopts "a:f:r:h" opt; do
    case $opt in
    a)
        Asm_Bundle_Version=$OPTARG
        ;;
    f)
        Flagger_Bundle_Version=$OPTARG
        ;;
    r)
        localRegistry=$OPTARG
        ;;
    h)
        usage
        ;;
    \?)
        echo "Invalid option: -$OPTARG" >&2
        usage
        ;;
    :)
        echo "Option -$OPTARG requires an argument." >&2
        usage
        ;;
    esac
done
check_bundle_version
load
push "${Asm_Bundle}" "${Asm_Bundle_Version}" "${localRegistry}"
push "${Flagger_Bundle}" "${Flagger_Bundle_Version}" "${localRegistry}"
proxyOn && sync
