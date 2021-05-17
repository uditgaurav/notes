#!/bin/bash
# This script is used to pull litmus image required to run generic experiments
# using litmus portal and push into your registry

set -e

list_all(){

declare -a portal_repository=("litmusportal-frontend" "litmusportal-server" "litmusportal-event-tracker"
                       "litmusportal-auth-server" "litmusportal-subscriber")
declare -a core_repository=("chaos-operator" "chaos-runner" "chaos-exporter" "go-runner")

if [[ -z "$CORE_TAG" ]];then
  CORE_TAG=$(get_latest_release)
fi

portal_tag=${1}

i=1
echo
for val in ${portal_repository[@]}; do
  echo "${i}. litmuschaos/${val}:${portal_tag}"
  i=$((i+1))
done
for val in ${core_repository[@]}; do
  echo "${i}. litmuschaos/${val}:${CORE_TAG}"
  i=$((i+1))
done
echo
printf "Other images are:
1. argoproj/workflow-controller:v2.9.3
2. mongo:4.2.8
3. argoproj/argocli:v2.9.3
4. argoproj/argoexec:v2.9.3

"
}

pull_all(){

declare -a portal_repository=("litmusportal-frontend" "litmusportal-server" "litmusportal-event-tracker"
                       "litmusportal-auth-server" "litmusportal-subscriber")
declare -a core_repository=("chaos-operator" "chaos-runner" "chaos-exporter" "go-runner")

if [[ -z "$CORE_TAG" ]];then
  export CORE_TAG=$(get_latest_release)
fi

local portal_tag=${1}

echo
for val in ${portal_repository[@]}; do
  echo " Pulling litmuschaos/${val}:${portal_tag}"
  docker pull litmuschaos/${val}:${portal_tag}
done
for val in ${core_repository[@]}; do
  echo " Pulling litmuschaos/${val}:${CORE_TAG}"
  docker pull litmuschaos/${val}:${CORE_TAG}
done
echo

declare -a other_images=("argoproj/workflow-controller:v2.9.3" "mongo:4.2.8" "argoproj/argocli:v2.9.3" "argoproj/argoexec:v2.9.3")
for val in ${other_images[@]}; do
  echo " Pulling ${val}"
  docker pull ${val}
done
echo
}

tag_and_push_all(){

echo "$portal_tag"

declare -a portal_repository=("litmusportal-frontend" "litmusportal-server" "litmusportal-event-tracker"
                       "litmusportal-auth-server" "litmusportal-subscriber")
declare -a core_repository=("chaos-operator" "chaos-runner" "chaos-exporter" "go-runner")

if [[ -z "$CORE_TAG" ]];then
  CORE_TAG=$(get_latest_release)
fi

local new_tag=${3}
local old_tag=${2}
local repo_name=${1}

echo
for val in ${portal_repository[@]}; do
  IMAGEID=$( docker images -q litmuschaos/${val}:${old_tag} )
  docker tag ${IMAGEID} ${repo_name}/${val}:${new_tag}
  docker push ${repo_name}/${val}:${new_tag}
done

for val in ${core_repository[@]}; do
  IMAGEID=$( docker images -q litmuschaos/${val}:${CORE_TAG} )
  docker tag ${IMAGEID} ${repo_name}/${val}:${new_tag}
  docker push ${repo_name}/${val}:${new_tag}
done
echo

declare -a other_images=("workflow-controller:v2.9.3" "argocli:v2.9.3" "argoexec:v2.9.3")
for val in ${other_images[@]}; do
  IMAGEID=$( docker images -q argoproj/${val} )
  docker tag ${IMAGEID} ${repo_name}/${val}
  docker push ${repo_name}/${val}
done
IMAGEID=$( docker images -q mongo:4.2.8 )
docker tag ${IMAGEID} ${repo_name}/mongo
docker push ${repo_name}/${val}
echo

}

get_latest_release() {
  curl --silent "https://api.github.com/repos/litmuschaos/litmus-go/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}



print_help(){
cat <<EOF

Usage: ${0} ARGS (list|pull|tag|push)

list:        "${0} list <image-tag>" will list all the images used by the litmus portal.     


pull:        "${0} pull <image-tag>" will pull the litmus images with the given image tag. 
              For example try running '${0} pull 2.0.0-Beta4', it will pull all the required 
              litmus image with tag 2.0.0-Beta4.
              For providing image tag for chore repositories (like chaos-exporter, chaos-runner,
              chaos-operator, litmus-go) you can export CORE_TAG=<image-tag> or it will fetch the
              latest image tag from latest release tag in the repo. 


push:          "${0} tag <repository> <old-image-tag> <new-image-tag>" will tag the litmus
               images with the given version and repository and push it. 
               For example try running '${0} uditgaurav 1.0' to tag the image with version 
               '1.0' and repository 'uditgaurav' and start pushing it.
               
EOF

}


case ${1} in
  list)
    list_all "${2}"
    ;;
  pull) 
    pull_all "${2}"
    ;;
  push) 
    tag_and_push_all "${2}" "${3}" "${4}"
    ;;
  *)
    print_help
    exit 1
esac
