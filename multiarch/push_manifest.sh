#!/bin/bash
# This is an experimental feature of docker.To enable experimental features in the Docker CLI, 
# edit the config.json (https://docs.docker.com/engine/reference/commandline/cli/#configuration-files)
# and set experimental to enabled. export DOCKER_CLI_EXPERIMENTAL=enabled
# Add "experimental": "enabled", to ~/.docker/config.json (default location) at the beginning of the file and not at the end.

setup(){

LITMUS_REPONAME="litmuschaos"
IMGNAME="go-runner"
IMGTAG="ci"
TARGET_REPONAME="uditgaurav"
TARGET_IMGNAME="go-runner"
TARGET_IMGTAG="test"

}

docker_login(){
    if [ ! -z ${DNAME} ] && [ ! -z ${DPASS} ];then
    echo "login to docker registry"
    echo ${DPASS} | docker login -u ${DNAME} --password-stdin
    fi
}

prepare_indivisual_manifest(){

   declare -ga platforms=("linux/arm64" "linux/amd64")

   for val in ${platforms[@]}; do
     arch=$(echo ${val} | cut -d / -f2)
     echo "pulling ${LITMUS_REPONAME}/${IMGNAME}:${IMGTAG} --platform ${val}"
     docker pull ${LITMUS_REPONAME}/${IMGNAME}:${IMGTAG} --platform ${val}
     IMAGEID=$( docker images -q ${LITMUS_REPONAME}/${IMGNAME}:${IMGTAG} )
     docker tag ${IMAGEID} ${TARGET_REPONAME}/${TARGET_IMGNAME}:${arch}
     echo "pushing ${TARGET_REPONAME}/${TARGET_IMGNAME}:${arch}"
     docker push ${TARGET_REPONAME}/${TARGET_IMGNAME}:${arch}
   done
}


create_and_retag_multiarch_manifest(){

   declare -ga platforms=("linux/arm64" "linux/amd64")

   cmd=""
   for val in ${platforms[@]}; do
        arch=$(echo ${val} | cut -d / -f2)
        cmd="${cmd} -a ${TARGET_REPONAME}/${TARGET_IMGNAME}:${arch}"
    done

    echo "Running: docker manifest create ${TARGET_REPONAME}/${TARGET_IMGNAME}:${TARGET_IMGTAG} ${cmd}"
    docker manifest create \
    ${TARGET_REPONAME}/${TARGET_IMGNAME}:${TARGET_IMGTAG} ${cmd}  
}

push_multiarch_manifest(){
    echo "pushing multiarch image"
    docker manifest push ${TARGET_REPONAME}/${TARGET_IMGNAME}:${TARGET_IMGTAG}
}


setup
docker_login
prepare_indivisual_manifest
create_and_retag_multiarch_manifest
push_multiarch_manifest
