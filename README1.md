## Table of content
- [Check the available commands](https://github.com/uditgaurav/docker-images#check-the-available-commands)
- [List Down all the Images](https://github.com/uditgaurav/docker-images#list-down-all-the-images)
- [Pull the LitmusChaos Images to your machine](https://github.com/uditgaurav/docker-images#pull-the-litmuschaos-images-to-your-machine)
- [Push the LitmusChaos Images to your repository](https://github.com/uditgaurav/docker-images#push-the-litmuschaos-images-to-your-repository)

## LitmusChaos Images

- This Repository contains all the images that are used to execute a litmuschaos generic experiment using litmus portal. For more information please check [LitmusChaos Repo](http://github.com/litmuschaos/litmus).


## Get LitmusChaos Images In Your Repository

- For pulling the litmus image and pushing into your registry please follow the given steps:

```bash
wget https://raw.githubusercontent.com/uditgaurav/docker-images/master/litmus_image_push.sh
chmod +x litmus_image_push.sh
```

#### Check the available commands:

```bash

udit@ubuntu ~/D/docker-images> bash litmus_image_push.sh -h

Usage: litmus_image_push.sh ARGS (list|pull|tag|push)

list:        "litmus_image_push.sh list <image-tag>" will list all the images used by the litmus portal.     


pull:        "litmus_image_push.sh pull <image-tag>" will pull the litmus images with the given image tag. 
              For example try running 'litmus_image_push.sh pull 2.0.0-Beta4', it will pull all the required 
              litmus image with tag 2.0.0-Beta4.
              For providing image tag for chore repositories (like chaos-exporter, chaos-runner,
              chaos-operator, litmus-go) you can export CORE_TAG=<image-tag> or it will fetch the
              latest image tag from latest release tag in the repo. 


push:          "litmus_image_push.sh tag <repository> <old-image-tag> <new-image-tag>" will tag the litmus
               images with the given version and repository and push it. 
               For example try running 'litmus_image_push.sh uditgaurav 1.0' to tag the image with version 
               '1.0' and repository 'uditgaurav' and start pushing it.

```

#### List Down all the Images

- For listing down the images you need to run the script with **TWO** first is list and second is the portal image tag to list as follow:

Format:
```
./litmus_image_push.sh list <IMAGE-TAG>
```

Example: 
```
./litmus_image_push.sh list 2.0.0-Beta4
```
```bash
root@ip-172-31-23-87:~# ./litmus_image_push.sh list 2.0.0-Beta4

1. litmuschaos/litmusportal-frontend:2.0.0-Beta4
2. litmuschaos/litmusportal-server:2.0.0-Beta4
3. litmuschaos/litmusportal-event-tracker:2.0.0-Beta4
4. litmuschaos/litmusportal-auth-server:2.0.0-Beta4
5. litmuschaos/litmusportal-subscriber:2.0.0-Beta4
6. litmuschaos/chaos-operator:1.13.3
7. litmuschaos/chaos-runner:1.13.3
8. litmuschaos/chaos-exporter:1.13.3
9. litmuschaos/go-runner:1.13.3

Other images are:
1. argoproj/workflow-controller:v2.9.3
2. mongo:4.2.8
3. argoproj/argocli:v2.9.3
4. argoproj/argoexec:v2.9.3
```

#### Pull the LitmusChaos Images to your machine

- For pulling LitmusChaos images to your machine run the script with **TWO** first one will be pull and second one will be the image tag you want to pull:

Format:
```
./litmus_image_push.sh pull <IMAGE-TAG>
```

Example: 
```
./litmus_image_push.sh pull 2.0.0-Beta4
```
```bash
root@ip-172-31-23-87:~# ./litmus_image_push.sh pull 2.0.0-Beta4

 Pulling litmuschaos/litmusportal-frontend:2.0.0-Beta4
2.0.0-Beta4: Pulling from litmuschaos/litmusportal-frontend
9b794450f7b6: Pull complete 
f8fd2f03c36e: Pull complete 
3cb08bcbe78b: Pull complete 
052722e881c1: Pull complete 
2fd682de3a2e: Pull complete 
80ce77084de4: Pull complete 
1b5bab11387e: Pull complete 
d64c83672bad: Pull complete 
6d9f710a197b: Pull complete 
5930ffa7e304: Pull complete 
Digest: sha256:509070409dd1abad59b0ea75607dd455f5157e0255552057910842b9d15f8ef6
Status: Downloaded newer image for litmuschaos/litmusportal-frontend:2.0.0-Beta4
docker.io/litmuschaos/litmusportal-frontend:2.0.0-Beta4
 Pulling litmuschaos/litmusportal-server:2.0.0-Beta4
2.0.0-Beta4: Pulling from litmuschaos/litmusportal-server
540db60ca938: Pull complete 
f061c865dd2a: Pull complete 
60a80321a714: Pull complete 
f2f214f06f1d: Pull complete 
Digest: sha256:f399cbea6979346498a431b37b1244255be9b70cae6d99275760c017853dbe78
Status: Downloaded newer image for litmuschaos/litmusportal-server:2.0.0-Beta4
docker.io/litmuschaos/litmusportal-server:2.0.0-Beta4
 Pulling litmuschaos/litmusportal-event-tracker:2.0.0-Beta4
2.0.0-Beta4: Pulling from litmuschaos/litmusportal-event-tracker
df20fa9351a1: Pull complete 
39d5d6854a1b: Pull complete 
52dff69fb26d: Pull complete 
Digest: sha256:c49d5c1ca4252484ce59ca13df82b669f7047701225f46b1e67a9bfe74a81696
Status: Downloaded newer image for litmuschaos/litmusportal-event-tracker:2.0.0-Beta4
docker.io/litmuschaos/litmusportal-event-tracker:2.0.0-Beta4
 Pulling litmuschaos/litmusportal-auth-server:2.0.0-Beta4
2.0.0-Beta4: Pulling from litmuschaos/litmusportal-auth-server
540db60ca938: Already exists 
f34d68243cc3: Pull complete 
34c2716b8c63: Pull complete 
Digest: sha256:22c63c5f1be9c4129011898a2b6c1cc0f942dba13ef2a15d2dc6803cacb91525
Status: Downloaded newer image for litmuschaos/litmusportal-auth-server:2.0.0-Beta4
docker.io/litmuschaos/litmusportal-auth-server:2.0.0-Beta4
 Pulling litmuschaos/litmusportal-subscriber:2.0.0-Beta4
2.0.0-Beta4: Pulling from litmuschaos/litmusportal-subscriber
df20fa9351a1: Already exists 
2709a9a996a6: Pull complete 
f799003e2cdd: Pull complete 
Digest: sha256:483838c5dbb1cac49f63332ae361e95a636484046413404d2bd625adbce8c500
Status: Downloaded newer image for litmuschaos/litmusportal-subscriber:2.0.0-Beta4
docker.io/litmuschaos/litmusportal-subscriber:2.0.0-Beta4
 Pulling litmuschaos/chaos-operator:1.13.3
1.13.3: Pulling from litmuschaos/chaos-operator
77a02d8cede1: Pull complete 
7777f1ac6191: Pull complete 
bcaa0be4cf95: Pull complete 
ba76012fd41e: Pull complete 
11962b97a2e6: Pull complete 
Digest: sha256:082fc3d4785b76a6f64a52cf73cf02c50dd5595790fabf284828d960130d1a4a
Status: Downloaded newer image for litmuschaos/chaos-operator:1.13.3
docker.io/litmuschaos/chaos-operator:1.13.3
 Pulling litmuschaos/chaos-runner:1.13.3
1.13.3: Pulling from litmuschaos/chaos-runner
77a02d8cede1: Already exists 
7777f1ac6191: Already exists 
4ed6ca33f4b3: Pull complete 
Digest: sha256:f840a61f177ee556d49dce7995725afa5689a1eb77582c1d06ae4b9c78205f2c
Status: Downloaded newer image for litmuschaos/chaos-runner:1.13.3
docker.io/litmuschaos/chaos-runner:1.13.3
 Pulling litmuschaos/chaos-exporter:1.13.3
1.13.3: Pulling from litmuschaos/chaos-exporter
540db60ca938: Already exists 
e409d31c2167: Pull complete 
709a22ae2cae: Pull complete 
Digest: sha256:8e851c1588f74915b104360fd284e1eacc6f535a38319dc792a12fe9db5b27a5
Status: Downloaded newer image for litmuschaos/chaos-exporter:1.13.3
docker.io/litmuschaos/chaos-exporter:1.13.3
 Pulling litmuschaos/go-runner:1.13.3
1.13.3: Pulling from litmuschaos/go-runner
540db60ca938: Already exists 
8f01f28f1540: Pull complete 
87db43b64cf7: Pull complete 
6b02da9ef052: Pull complete 
60d80da56898: Pull complete 
660832c79af5: Pull complete 
a2fc46b202e3: Pull complete 
65ac9b37fd35: Pull complete 
b921de47829b: Pull complete 
ce6443a7987f: Pull complete 
94d9cbcff94a: Pull complete 
2fcaaf006032: Pull complete 
fd71b6108a1d: Pull complete 
Digest: sha256:bc0ecc4872725248cd709fd0969dca25704b91686e8048a94ae0c8caa27150b9
Status: Downloaded newer image for litmuschaos/go-runner:1.13.3
docker.io/litmuschaos/go-runner:1.13.3

 Pulling argoproj/workflow-controller:v2.9.3
v2.9.3: Pulling from argoproj/workflow-controller
1f407d3f644c: Pull complete 
1633c91701ba: Pull complete 
Digest: sha256:9f518c4c366121ec252ec4941148aecf3f95e8f18f3ab1e8d44788038ad794c0
Status: Downloaded newer image for argoproj/workflow-controller:v2.9.3
docker.io/argoproj/workflow-controller:v2.9.3
 Pulling mongo:4.2.8
4.2.8: Pulling from library/mongo
f08d8e2a3ba1: Pull complete 
3baa9cb2483b: Pull complete 
94e5ff4c0b15: Pull complete 
1860925334f9: Pull complete 
9d42806c06e6: Pull complete 
31a9fd218257: Pull complete 
5bd6e3f73ab9: Pull complete 
f6ae7a64936b: Pull complete 
a614d629c284: Pull complete 
477320af2dcc: Pull complete 
b8aab702fcf5: Pull complete 
b94c6a2dc294: Pull complete 
8cf889bdb7c6: Pull complete 
Digest: sha256:bb3616e2f78a37bb607f59e451922d63d27e81d272514fb1cbffc7d1eab00eaf
Status: Downloaded newer image for mongo:4.2.8
docker.io/library/mongo:4.2.8
 Pulling argoproj/argocli:v2.9.3
v2.9.3: Pulling from argoproj/argocli
a2f8c580559a: Pull complete 
e7f20eeaa3a5: Pull complete 
06461e61bd4f: Pull complete 
79097b56e3b3: Pull complete 
aef4bbcef25f: Pull complete 
Digest: sha256:4353763dc8d748a6e922c81695af6b92524691053df9e6dc19618c2e5bfa7832
Status: Downloaded newer image for argoproj/argocli:v2.9.3
docker.io/argoproj/argocli:v2.9.3
 Pulling argoproj/argoexec:v2.9.3
v2.9.3: Pulling from argoproj/argoexec
54fec2fa59d0: Pull complete 
ea1b58a62c8f: Pull complete 
baf5dcb8fbf2: Pull complete 
f7d968f2f223: Pull complete 
7be27d15fd7b: Pull complete 
1648478b409a: Pull complete 
220fc1c30111: Pull complete 
f102822e7b4e: Pull complete 
2ce8271a1891: Pull complete 
Digest: sha256:a25cbe11d210b67e9621fecd4505e3438d0dd3ce8fc5e8d3dea1bbc2adb853f6
Status: Downloaded newer image for argoproj/argoexec:v2.9.3
docker.io/argoproj/argoexec:v2.9.3
```

#### Push the LitmusChaos Images to your repository

- For pushing the pulled LitmusChaos images into your repository run the following command:

Format:
```
./litmus_image_push.sh push <YOUR-REPOSITORY-NAME> <OLD-IMAGE-TAG> <NEW-IMAGE-TAG>
```

Example:
```
./litmus_image_push.sh push uditgaurav 2.0.0-Beta4 2.0.0-Beta4
```

```bash
root@ip-172-31-23-87:~# bash litmus_image_push.sh push uditgaurav 2.0.0-Beta4 2.0.0-Beta4


The push refers to repository [docker.io/uditgaurav/litmusportal-frontend]
1883b233390f: Layer already exists 
59607d089ef7: Layer already exists 
7457bd802d98: Layer already exists 
db639965f176: Layer already exists 
19e5eaa18644: Layer already exists 
697f2aa6662e: Layer already exists 
1f9e2810747e: Layer already exists 
a3355a4d5656: Layer already exists 
50a03d8e0394: Layer already exists 
2b2bcc6e6724: Layer already exists 
2.0.0-Beta4: digest: sha256:509070409dd1abad59b0ea75607dd455f5157e0255552057910842b9d15f8ef6 size: 2403
The push refers to repository [docker.io/uditgaurav/litmusportal-server]
e7000e2a5eee: Layer already exists 
fff11f3178f9: Layer already exists 
885e33cc66f3: Layer already exists 
b2d5eeeaba3a: Layer already exists 
2.0.0-Beta4: digest: sha256:f399cbea6979346498a431b37b1244255be9b70cae6d99275760c017853dbe78 size: 1156
The push refers to repository [docker.io/uditgaurav/litmusportal-event-tracker]
5b3ef05529ee: Layer already exists 
dc76fe0fc4f5: Layer already exists 
50644c29ef5a: Layer already exists 
2.0.0-Beta4: digest: sha256:c49d5c1ca4252484ce59ca13df82b669f7047701225f46b1e67a9bfe74a81696 size: 948
The push refers to repository [docker.io/uditgaurav/litmusportal-auth-server]
bfc9c1cd14e1: Layer already exists 
4357ce4eab52: Layer already exists 
b2d5eeeaba3a: Layer already exists 
2.0.0-Beta4: digest: sha256:22c63c5f1be9c4129011898a2b6c1cc0f942dba13ef2a15d2dc6803cacb91525 size: 947
The push refers to repository [docker.io/uditgaurav/litmusportal-subscriber]
99ae9f9eab71: Layer already exists 
f200ecc80756: Layer already exists 
50644c29ef5a: Layer already exists 
2.0.0-Beta4: digest: sha256:483838c5dbb1cac49f63332ae361e95a636484046413404d2bd625adbce8c500 size: 948
The push refers to repository [docker.io/uditgaurav/chaos-operator]
62ee58bf5f77: Layer already exists 
906cffa1729a: Layer already exists 
87e059e96c9d: Layer already exists 
8838b2c54cd7: Layer already exists 
41963ce1cb78: Layer already exists 
2.0.0-Beta4: digest: sha256:1867710d22474afee52f1cb2aac93da2e87fbd6a6da8d68d29e7cc9fa6da1a77 size: 1363
The push refers to repository [docker.io/uditgaurav/chaos-runner]
1a41be1519e6: Layer already exists 
8838b2c54cd7: Layer already exists 
41963ce1cb78: Layer already exists 
2.0.0-Beta4: digest: sha256:61c9eca851cd548677e35302642f32120fe152c3507dd60e92860f649a746257 size: 949
The push refers to repository [docker.io/uditgaurav/chaos-exporter]
378a50850872: Layer already exists 
fe86cb641cac: Layer already exists 
b2d5eeeaba3a: Layer already exists 
2.0.0-Beta4: digest: sha256:0d1a325c61fd57c45bb92836010a3bcc1b4e46f93416559c5926ea4d80b833a9 size: 948
The push refers to repository [docker.io/uditgaurav/go-runner]
2f8370cf981f: Layer already exists 
1694355f7c8e: Layer already exists 
431afeadf93b: Layer already exists 
662db7a855f5: Layer already exists 
24682565c58c: Layer already exists 
9025cb22bcf3: Layer already exists 
9640c560222b: Layer already exists 
c4c828e84c59: Layer already exists 
409cc9fd343a: Layer already exists 
bf3bd1ddeb30: Layer already exists 
11bfb6ecefea: Layer already exists 
4ce9da4ee628: Layer already exists 
b2d5eeeaba3a: Layer already exists 
2.0.0-Beta4: digest: sha256:16d50fa04c465976b6c356e687fb84740a22c5ad5c88ace785ae44a75da7d598 size: 3058

The push refers to repository [docker.io/uditgaurav/workflow-controller]
a199faf3f870: Mounted from argoproj/workflow-controller 
ec0a2776976b: Mounted from argoproj/workflow-controller 
v2.9.3: digest: sha256:9f518c4c366121ec252ec4941148aecf3f95e8f18f3ab1e8d44788038ad794c0 size: 738
The push refers to repository [docker.io/uditgaurav/argocli]
95121346ef9a: Mounted from argoproj/argocli 
fbbc9cc4e41d: Mounted from argoproj/argocli 
c5ed1c17bd56: Mounted from argoproj/argocli 
0378eb7d8202: Mounted from argoproj/argocli 
82ef510fb22f: Mounted from argoproj/argocli 
v2.9.3: digest: sha256:4353763dc8d748a6e922c81695af6b92524691053df9e6dc19618c2e5bfa7832 size: 1363
The push refers to repository [docker.io/uditgaurav/argoexec]
ed60a3b92d3c: Mounted from argoproj/argoexec 
75515e229396: Mounted from argoproj/argoexec 
9f67644d8e91: Mounted from argoproj/argoexec 
5906eba2497c: Mounted from argoproj/argoexec 
b2fe9b4eb3ed: Mounted from argoproj/argoexec 
71d40540ea58: Mounted from argoproj/argoexec 
22d1b6b8a5c2: Mounted from argoproj/argoexec 
f2b050ecb00c: Mounted from argoproj/argoexec 
c2adabaecedb: Mounted from argoproj/argoexec 
v2.9.3: digest: sha256:a25cbe11d210b67e9621fecd4505e3438d0dd3ce8fc5e8d3dea1bbc2adb853f6 size: 2209
The push refers to repository [docker.io/uditgaurav/argoexec]
ed60a3b92d3c: Layer already exists 
75515e229396: Layer already exists 
9f67644d8e91: Layer already exists 
5906eba2497c: Layer already exists 
b2fe9b4eb3ed: Layer already exists 
71d40540ea58: Layer already exists 
22d1b6b8a5c2: Layer already exists 
f2b050ecb00c: Layer already exists 
c2adabaecedb: Layer already exists 
v2.9.3: digest: sha256:a25cbe11d210b67e9621fecd4505e3438d0dd3ce8fc5e8d3dea1bbc2adb853f6
```
