# kubectl
[官方文档kubectl](https://kubernetes.io/docs/reference/kubectl/)

kubectl 的命令格式
```text
kubectl [command] [TYPE] [NAME] [flags]
```
command 表示操作 比如get、describe、create、delete、apply ....
type 表示资源类型 比如 pod 、 deployment 、replicaset 、logs ....
name 表示资源的名称信息
flags 表示一下特性，比如 -o 表示输出


进入镜像容器的控制台
```bash
kubectl exec pod_name -n namespace -i -t -- bash
```


获取指定namespace下的所有的deployment 名称
```bash
kubectl get deployments -n tce

kubectl get deployments -n tce | grep "amp"
```

获取指定deployment 名称的yaml信息
```bash
kubectl get deployment tcloud-barad-alarm-amp -n tce -o yaml > 1.yaml
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "91"
    infra.tce.io/last-applied-configuration: '{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{"infra.tce.io/last-applied-definition":"8dbfcc8af723bb71c6d4262f61ede147"},"creationTimestamp":null,"labels":{"infra.tce.io/app-version":"3.10.11-20250214-160805-a57bd60.rhel.amd64","infra.tce.io/chart-name":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","infra.tce.io/oam-product":"barad"},"name":"tcloud-barad-alarm-amp","namespace":"tce","ownerReferences":[{"apiVersion":"infra.tce.io/v1","blockOwnerDeletion":true,"controller":true,"kind":"Application","name":"tcloud-barad-alarm-amp","uid":"11f1d998-cc7f-42a7-bb8d-e66a3263851d"}]},"spec":{"progressDeadlineSeconds":600,"replicas":2,"revisionHistoryLimit":10,"selector":{"matchLabels":{"app":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"strategy":{"rollingUpdate":{"maxSurge":"25%","maxUnavailable":"25%"},"type":"RollingUpdate"},"template":{"metadata":{"creationTimestamp":null,"labels":{"app":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"spec":{"affinity":{"podAntiAffinity":{"preferredDuringSchedulingIgnoredDuringExecution":[{"podAffinityTerm":{"labelSelector":{"matchExpressions":[{"key":"app","operator":"In","values":["tcloud-barad-alarm-amp"]}]},"topologyKey":"kubernetes.io/hostname"},"weight":100}]}},"containers":[{"command":["/bin/bash","-c","/usr/bin/python
      /usr/bin/supervisord -c /etc/supervisord.conf \u0026\u0026 /usr/sbin/crond \u0026\u0026
      tailf /etc/hosts"],"env":[{"name":"APPLICATION_NAME","value":"tcloud-barad-alarm-amp"},{"name":"MY_POD_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.podIP"}}},{"name":"MY_POD_HOSTIP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.hostIP"}}},{"name":"MY_POD_NAME","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.name"}}}],"image":"registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64","imagePullPolicy":"IfNotPresent","livenessProbe":{"exec":{"command":["bash","/tce/healthchk.sh"]},"failureThreshold":3,"initialDelaySeconds":30,"periodSeconds":20,"successThreshold":1,"timeoutSeconds":5},"name":"tcloud-barad-alarm-amp","ports":[{"containerPort":9411,"name":"port-0"}],"readinessProbe":{"exec":{"command":["bash","/tce/healthchk.sh"]},"failureThreshold":3,"initialDelaySeconds":30,"periodSeconds":20,"successThreshold":1,"timeoutSeconds":5},"resources":{"limits":{"cpu":"2","memory":"4G"},"requests":{"cpu":"1","memory":"2G"}},"volumeMounts":[{"mountPath":"/tce/conf/global","name":"volume-config"},{"mountPath":"/tce/conf/cm","name":"volume-local-config"},{"mountPath":"/data/storage","name":"volume-log"},{"mountPath":"/tce/customize","name":"volume-customize-confd","readOnly":true},{"mountPath":"/etc/supervisord.d","name":"volume-supervisord-confd","readOnly":true},{"mountPath":"/sys/fs/cgroup","name":"volume-centos-cgroup","readOnly":true},{"mountPath":"/etc/localtime","name":"volume-zoneinfo","readOnly":true},{"mountPath":"/data/customize_packet","name":"volume-customize"},{"mountPath":"/data/tce.config.center","name":"volume-host-path-0","readOnly":true}]}],"dnsConfig":{"options":[{"name":"single-request-reopen"},{"name":"timeout","value":"1"}]},"dnsPolicy":"ClusterFirst","restartPolicy":"Always","schedulerName":"default-scheduler","securityContext":{},"terminationGracePeriodSeconds":30,"topologySpreadConstraints":[{"labelSelector":{"matchLabels":{"app":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"maxSkew":1,"topologyKey":"topology.kubernetes.io/zone","whenUnsatisfiable":"DoNotSchedule"}],"volumes":[{"configMap":{"items":[{"key":"global.json","mode":0,"path":"global.json"}],"name":"global.conf.d"},"name":"volume-config"},{"configMap":{"items":[{"key":"local.json","mode":0,"path":"local.json"}],"name":"tcloud-barad-alarm-amp"},"name":"volume-local-config"},{"hostPath":{"path":"/data/k8s/log/tce/tcloud-barad-alarm-amp"},"name":"volume-log"},{"configMap":{"items":[{"key":"customize_deploy.sh","mode":0,"path":"customize_deploy.sh"},{"key":"customize_move.py","mode":0,"path":"customize_move.py"}],"name":"global.customize.d"},"name":"volume-customize-confd"},{"configMap":{"items":[{"key":"supervisord.conf","mode":0,"path":"supervisord.default.ini"}],"name":"global.conf.d"},"name":"volume-supervisord-confd"},{"hostPath":{"path":"/sys/fs/cgroup"},"name":"volume-centos-cgroup"},{"hostPath":{"path":"/etc/localtime"},"name":"volume-zoneinfo"},{"hostPath":{"path":"/data/customize_packet"},"name":"volume-customize"},{"hostPath":{"path":"/data/tce.config.center"},"name":"volume-host-path-0"}]}}},"status":{}}'
    infra.tce.io/last-applied-definition: 8dbfcc8af723bb71c6d4262f61ede147
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{"deployment.kubernetes.io/revision":"89","infra.tce.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{\"infra.tce.io/last-applied-definition\":\"8dbfcc8af723bb71c6d4262f61ede147\"},\"creationTimestamp\":null,\"labels\":{\"infra.tce.io/app-version\":\"3.10.11-20250214-160805-a57bd60.rhel.amd64\",\"infra.tce.io/chart-name\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-app\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-comp\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-product\":\"barad\"},\"name\":\"tcloud-barad-alarm-amp\",\"namespace\":\"tce\",\"ownerReferences\":[{\"apiVersion\":\"infra.tce.io/v1\",\"blockOwnerDeletion\":true,\"controller\":true,\"kind\":\"Application\",\"name\":\"tcloud-barad-alarm-amp\",\"uid\":\"11f1d998-cc7f-42a7-bb8d-e66a3263851d\"}]},\"spec\":{\"progressDeadlineSeconds\":600,\"replicas\":2,\"revisionHistoryLimit\":10,\"selector\":{\"matchLabels\":{\"app\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-app\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-comp\":\"tcloud-barad-alarm-amp\",\"module\":\"tcloud-barad\"}},\"strategy\":{\"rollingUpdate\":{\"maxSurge\":\"25%\",\"maxUnavailable\":\"25%\"},\"type\":\"RollingUpdate\"},\"template\":{\"metadata\":{\"creationTimestamp\":null,\"labels\":{\"app\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-app\":\"tcloud-barad-alarm-amp\",\"infra.tce.io/oam-comp\":\"tcloud-barad-alarm-amp\",\"module\":\"tcloud-barad\"}},\"spec\":{\"affinity\":{\"podAntiAffinity\":{\"preferredDuringSchedulingIgnoredDuringExecution\":[{\"podAffinityTerm\":{\"labelSelector\":{\"matchExpressions\":[{\"key\":\"app\",\"operator\":\"In\",\"values\":[\"tcloud-barad-alarm-amp\"]}]},\"topologyKey\":\"kubernetes.io/hostname\"},\"weight\":100}]}},\"containers\":[{\"command\":[\"/bin/bash\",\"-c\",\"/usr/bin/python /usr/bin/supervisord -c /etc/supervisord.conf \\u0026\\u0026 /usr/sbin/crond \\u0026\\u0026 tailf /etc/hosts\"],\"env\":[{\"name\":\"APPLICATION_NAME\",\"value\":\"tcloud-barad-alarm-amp\"},{\"name\":\"MY_POD_IP\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"status.podIP\"}}},{\"name\":\"MY_POD_HOSTIP\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"status.hostIP\"}}},{\"name\":\"MY_POD_NAME\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"metadata.name\"}}}],\"image\":\"registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64\",\"imagePullPolicy\":\"IfNotPresent\",\"livenessProbe\":{\"exec\":{\"command\":[\"bash\",\"/tce/healthchk.sh\"]},\"failureThreshold\":3,\"initialDelaySeconds\":30,\"periodSeconds\":20,\"successThreshold\":1,\"timeoutSeconds\":5},\"name\":\"tcloud-barad-alarm-amp\",\"ports\":[{\"containerPort\":9411,\"name\":\"port-0\"}],\"readinessProbe\":{\"exec\":{\"command\":[\"bash\",\"/tce/healthchk.sh\"]},\"failureThreshold\":3,\"initialDelaySeconds\":30,\"periodSeconds\":20,\"successThreshold\":1,\"timeoutSeconds\":5},\"resources\":{\"limits\":{\"cpu\":\"2\",\"memory\":\"4G\"},\"requests\":{\"cpu\":\"1\",\"memory\":\"2G\"}},\"volumeMounts\":[{\"mountPath\":\"/tce/conf/global\",\"name\":\"volume-config\"},{\"mountPath\":\"/tce/conf/cm\",\"name\":\"volume-local-config\"},{\"mountPath\":\"/data/storage\",\"name\":\"volume-log\"},{\"mountPath\":\"/tce/customize\",\"name\":\"volume-customize-confd\",\"readOnly\":true},{\"mountPath\":\"/etc/supervisord.d\",\"name\":\"volume-supervisord-confd\",\"readOnly\":true},{\"mountPath\":\"/sys/fs/cgroup\",\"name\":\"volume-centos-cgroup\",\"readOnly\":true},{\"mountPath\":\"/etc/localtime\",\"name\":\"volume-zoneinfo\",\"readOnly\":true},{\"mountPath\":\"/data/customize_packet\",\"name\":\"volume-customize\"},{\"mountPath\":\"/data/tce.config.center\",\"name\":\"volume-host-path-0\",\"readOnly\":true}]}],\"dnsConfig\":{\"options\":[{\"name\":\"single-request-reopen\"},{\"name\":\"timeout\",\"value\":\"1\"}]},\"dnsPolicy\":\"ClusterFirst\",\"restartPolicy\":\"Always\",\"schedulerName\":\"default-scheduler\",\"securityContext\":{},\"terminationGracePeriodSeconds\":30,\"topologySpreadConstraints\":[{\"labelSelector\":{\"matchLabels\":{\"app\":\"tcloud-barad-alarm-amp\",\"module\":\"tcloud-barad\"}},\"maxSkew\":1,\"topologyKey\":\"topology.kubernetes.io/zone\",\"whenUnsatisfiable\":\"DoNotSchedule\"}],\"volumes\":[{\"configMap\":{\"items\":[{\"key\":\"global.json\",\"mode\":0,\"path\":\"global.json\"}],\"name\":\"global.conf.d\"},\"name\":\"volume-config\"},{\"configMap\":{\"items\":[{\"key\":\"local.json\",\"mode\":0,\"path\":\"local.json\"}],\"name\":\"tcloud-barad-alarm-amp\"},\"name\":\"volume-local-config\"},{\"hostPath\":{\"path\":\"/data/k8s/log/tce/tcloud-barad-alarm-amp\"},\"name\":\"volume-log\"},{\"configMap\":{\"items\":[{\"key\":\"customize_deploy.sh\",\"mode\":0,\"path\":\"customize_deploy.sh\"},{\"key\":\"customize_move.py\",\"mode\":0,\"path\":\"customize_move.py\"}],\"name\":\"global.customize.d\"},\"name\":\"volume-customize-confd\"},{\"configMap\":{\"items\":[{\"key\":\"supervisord.conf\",\"mode\":0,\"path\":\"supervisord.default.ini\"}],\"name\":\"global.conf.d\"},\"name\":\"volume-supervisord-confd\"},{\"hostPath\":{\"path\":\"/sys/fs/cgroup\"},\"name\":\"volume-centos-cgroup\"},{\"hostPath\":{\"path\":\"/etc/localtime\"},\"name\":\"volume-zoneinfo\"},{\"hostPath\":{\"path\":\"/data/customize_packet\"},\"name\":\"volume-customize\"},{\"hostPath\":{\"path\":\"/data/tce.config.center\"},\"name\":\"volume-host-path-0\"}]}}},\"status\":{}}","infra.tce.io/last-applied-definition":"8dbfcc8af723bb71c6d4262f61ede147"},"creationTimestamp":"2024-07-17T07:48:38Z","generation":107,"labels":{"infra.tce.io/app-version":"3.10.11-20250214-160805-a57bd60.rhel.amd64","infra.tce.io/chart-name":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","infra.tce.io/oam-product":"barad"},"name":"tcloud-barad-alarm-amp","namespace":"tce","ownerReferences":[{"apiVersion":"infra.tce.io/v1","blockOwnerDeletion":true,"controller":true,"kind":"Application","name":"tcloud-barad-alarm-amp","uid":"11f1d998-cc7f-42a7-bb8d-e66a3263851d"}],"resourceVersion":"2117179151","selfLink":"/apis/apps/v1/namespaces/tce/deployments/tcloud-barad-alarm-amp","uid":"5b3e8757-1b5f-4728-aad5-e3653bcae0c8"},"spec":{"progressDeadlineSeconds":600,"replicas":2,"revisionHistoryLimit":10,"selector":{"matchLabels":{"app":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"strategy":{"rollingUpdate":{"maxSurge":"25%","maxUnavailable":"25%"},"type":"RollingUpdate"},"template":{"metadata":{"creationTimestamp":null,"labels":{"app":"tcloud-barad-alarm-amp","infra.tce.io/oam-app":"tcloud-barad-alarm-amp","infra.tce.io/oam-comp":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"spec":{"affinity":{"podAntiAffinity":{"preferredDuringSchedulingIgnoredDuringExecution":[{"podAffinityTerm":{"labelSelector":{"matchExpressions":[{"key":"app","operator":"In","values":["tcloud-barad-alarm-amp"]}]},"topologyKey":"kubernetes.io/hostname"},"weight":100}]}},"containers":[{"command":["/bin/bash","-c","/usr/bin/python /usr/bin/supervisord -c /etc/supervisord.conf \u0026\u0026 /usr/sbin/crond \u0026\u0026 tailf /etc/hosts"],"env":[{"name":"TZ","value":"Europe/Moscow"},{"name":"APPLICATION_NAME","value":"tcloud-barad-alarm-amp"},{"name":"MY_POD_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.podIP"}}},{"name":"MY_POD_HOSTIP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.hostIP"}}},{"name":"MY_POD_NAME","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.name"}}}],"image":"registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64","imagePullPolicy":"IfNotPresent","livenessProbe":{"exec":{"command":["bash","/tce/healthchk.sh"]},"failureThreshold":3,"initialDelaySeconds":30,"periodSeconds":20,"successThreshold":1,"timeoutSeconds":5},"name":"tcloud-barad-alarm-amp","ports":[{"containerPort":9411,"name":"port-0","protocol":"TCP"}],"readinessProbe":{"exec":{"command":["bash","/tce/healthchk.sh"]},"failureThreshold":3,"initialDelaySeconds":30,"periodSeconds":20,"successThreshold":1,"timeoutSeconds":5},"resources":{"limits":{"cpu":"2","memory":"4G"},"requests":{"cpu":"1","memory":"2G"}},"terminationMessagePath":"/dev/termination-log","terminationMessagePolicy":"File","volumeMounts":[{"mountPath":"/tce/conf/global","name":"volume-config"},{"mountPath":"/tce/conf/cm","name":"volume-local-config"},{"mountPath":"/data/storage","name":"volume-log"},{"mountPath":"/tce/customize","name":"volume-customize-confd","readOnly":true},{"mountPath":"/etc/supervisord.d","name":"volume-supervisord-confd","readOnly":true},{"mountPath":"/sys/fs/cgroup","name":"volume-centos-cgroup","readOnly":true},{"mountPath":"/etc/localtime","name":"volume-zoneinfo","readOnly":true},{"mountPath":"/data/customize_packet","name":"volume-customize"},{"mountPath":"/data/tce.config.center","name":"volume-host-path-0","readOnly":true}]}],"dnsConfig":{"options":[{"name":"single-request-reopen"},{"name":"timeout","value":"1"}]},"dnsPolicy":"ClusterFirst","restartPolicy":"Always","schedulerName":"default-scheduler","securityContext":{},"terminationGracePeriodSeconds":30,"topologySpreadConstraints":[{"labelSelector":{"matchLabels":{"app":"tcloud-barad-alarm-amp","module":"tcloud-barad"}},"maxSkew":1,"topologyKey":"topology.kubernetes.io/zone","whenUnsatisfiable":"DoNotSchedule"}],"volumes":[{"configMap":{"defaultMode":420,"items":[{"key":"global.json","mode":0,"path":"global.json"}],"name":"global.conf.d"},"name":"volume-config"},{"configMap":{"defaultMode":420,"items":[{"key":"local.json","mode":0,"path":"local.json"}],"name":"tcloud-barad-alarm-amp"},"name":"volume-local-config"},{"hostPath":{"path":"/data/k8s/log/tce/tcloud-barad-alarm-amp","type":""},"name":"volume-log"},{"configMap":{"defaultMode":420,"items":[{"key":"customize_deploy.sh","mode":0,"path":"customize_deploy.sh"},{"key":"customize_move.py","mode":0,"path":"customize_move.py"}],"name":"global.customize.d"},"name":"volume-customize-confd"},{"configMap":{"defaultMode":420,"items":[{"key":"supervisord.conf","mode":0,"path":"supervisord.default.ini"}],"name":"global.conf.d"},"name":"volume-supervisord-confd"},{"hostPath":{"path":"/sys/fs/cgroup","type":""},"name":"volume-centos-cgroup"},{"hostPath":{"path":"/etc/localtime","type":""},"name":"volume-zoneinfo"},{"hostPath":{"path":"/data/customize_packet","type":""},"name":"volume-customize"},{"hostPath":{"path":"/data/tce.config.center","type":""},"name":"volume-host-path-0"}]}}},"status":{"availableReplicas":2,"conditions":[{"lastTransitionTime":"2025-02-14T01:38:27Z","lastUpdateTime":"2025-02-14T01:38:27Z","message":"Deployment has minimum availability.","reason":"MinimumReplicasAvailable","status":"True","type":"Available"},{"lastTransitionTime":"2025-02-07T04:05:13Z","lastUpdateTime":"2025-02-14T08:16:24Z","message":"ReplicaSet \"tcloud-barad-alarm-amp-7cb448b4c4\" has successfully progressed.","reason":"NewReplicaSetAvailable","status":"True","type":"Progressing"}],"observedGeneration":107,"readyReplicas":2,"replicas":2,"updatedReplicas":2}}
  creationTimestamp: "2024-07-17T07:48:38Z"
  generation: 108
  labels:
    infra.tce.io/app-version: 3.10.11-20250214-160805-a57bd60.rhel.amd64
    infra.tce.io/chart-name: tcloud-barad-alarm-amp
    infra.tce.io/oam-app: tcloud-barad-alarm-amp
    infra.tce.io/oam-comp: tcloud-barad-alarm-amp
    infra.tce.io/oam-product: barad
  name: tcloud-barad-alarm-amp
  namespace: tce
  ownerReferences:
  - apiVersion: infra.tce.io/v1
    blockOwnerDeletion: true
    controller: true
    kind: Application
    name: tcloud-barad-alarm-amp
    uid: 11f1d998-cc7f-42a7-bb8d-e66a3263851d
  resourceVersion: "2117684741"
  selfLink: /apis/apps/v1/namespaces/tce/deployments/tcloud-barad-alarm-amp
  uid: 5b3e8757-1b5f-4728-aad5-e3653bcae0c8
spec:
  progressDeadlineSeconds: 600
  replicas: 2
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: tcloud-barad-alarm-amp
      infra.tce.io/oam-app: tcloud-barad-alarm-amp
      infra.tce.io/oam-comp: tcloud-barad-alarm-amp
      module: tcloud-barad
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: tcloud-barad-alarm-amp
        infra.tce.io/oam-app: tcloud-barad-alarm-amp
        infra.tce.io/oam-comp: tcloud-barad-alarm-amp
        module: tcloud-barad
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - tcloud-barad-alarm-amp
              topologyKey: kubernetes.io/hostname
            weight: 100
      containers:
      - command:
        - /bin/bash
        - -c
        - /usr/bin/python /usr/bin/supervisord -c /etc/supervisord.conf && /usr/sbin/crond
          && tailf /etc/hosts
        env:
        - name: TZ
          value: Asia/Shanghai
        - name: APPLICATION_NAME
          value: tcloud-barad-alarm-amp
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: status.podIP
        - name: MY_POD_HOSTIP
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: status.hostIP
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        image: registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64
        imagePullPolicy: IfNotPresent
        livenessProbe:
          exec:
            command:
            - bash
            - /tce/healthchk.sh
          failureThreshold: 3
          initialDelaySeconds: 30
          periodSeconds: 20
          successThreshold: 1
          timeoutSeconds: 5
        name: tcloud-barad-alarm-amp
        ports:
        - containerPort: 9411
          name: port-0
          protocol: TCP
        readinessProbe:
          exec:
            command:
            - bash
            - /tce/healthchk.sh
          failureThreshold: 3
          initialDelaySeconds: 30
          periodSeconds: 20
          successThreshold: 1
          timeoutSeconds: 5
        resources:
          limits:
            cpu: "2"
            memory: 4G
          requests:
            cpu: "1"
            memory: 2G
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /tce/conf/global
          name: volume-config
        - mountPath: /tce/conf/cm
          name: volume-local-config
        - mountPath: /data/storage
          name: volume-log
        - mountPath: /tce/customize
          name: volume-customize-confd
          readOnly: true
        - mountPath: /etc/supervisord.d
          name: volume-supervisord-confd
          readOnly: true
        - mountPath: /sys/fs/cgroup
          name: volume-centos-cgroup
          readOnly: true
        - mountPath: /etc/localtime
          name: volume-zoneinfo
          readOnly: true
        - mountPath: /data/customize_packet
          name: volume-customize
        - mountPath: /data/tce.config.center
          name: volume-host-path-0
          readOnly: true
      dnsConfig:
        options:
        - name: single-request-reopen
        - name: timeout
          value: "1"
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
      topologySpreadConstraints:
      - labelSelector:
          matchLabels:
            app: tcloud-barad-alarm-amp
            module: tcloud-barad
        maxSkew: 1
        topologyKey: topology.kubernetes.io/zone
        whenUnsatisfiable: DoNotSchedule
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: global.json
            mode: 0
            path: global.json
          name: global.conf.d
        name: volume-config
      - configMap:
          defaultMode: 420
          items:
          - key: local.json
            mode: 0
            path: local.json
          name: tcloud-barad-alarm-amp
        name: volume-local-config
      - hostPath:
          path: /data/k8s/log/tce/tcloud-barad-alarm-amp
          type: ""
        name: volume-log
      - configMap:
          defaultMode: 420
          items:
          - key: customize_deploy.sh
            mode: 0
            path: customize_deploy.sh
          - key: customize_move.py
            mode: 0
            path: customize_move.py
          name: global.customize.d
        name: volume-customize-confd
      - configMap:
          defaultMode: 420
          items:
          - key: supervisord.conf
            mode: 0
            path: supervisord.default.ini
          name: global.conf.d
        name: volume-supervisord-confd
      - hostPath:
          path: /sys/fs/cgroup
          type: ""
        name: volume-centos-cgroup
      - hostPath:
          path: /etc/localtime
          type: ""
        name: volume-zoneinfo
      - hostPath:
          path: /data/customize_packet
          type: ""
        name: volume-customize
      - hostPath:
          path: /data/tce.config.center
          type: ""
        name: volume-host-path-0
status:
  availableReplicas: 2
  collisionCount: 2
  conditions:
  - lastTransitionTime: "2025-02-14T09:03:00Z"
    lastUpdateTime: "2025-02-14T09:03:00Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2025-02-07T04:05:13Z"
    lastUpdateTime: "2025-02-14T09:34:45Z"
    message: ReplicaSet "tcloud-barad-alarm-amp-57449d4fd4" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 108
  readyReplicas: 2
  replicas: 2
  updatedReplicas: 2
```


获取指定 namespace 下所有的 replicaset 名称
```bash

kubectl get replicasets -n tce

// 这里一般会查询到很多的,因为基于replicaset创建的pod每次都会生成一个
kubectl get replicasets -n tce | grep 'amp' 
```

获取指定的replicaset名称的yaml信息
```bash
kubectl get replicaset replicaset-name -n tce -o yaml > replicaset.yaml
```
```yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    ipip.ipv4.network.infra.tce.io/address: 172.16.13.71
    ipip.ipv4.network.infra.tce.io/attributes: "null"
    network.infra.tce.io/ipam-allocation-error: ""
    network.infra.tce.io/ipv4: 172.16.13.71
    network.infra.tce.io/type: ipip
    v1.multus-cni.io/default-network: ipip
  creationTimestamp: "2025-02-14T09:45:38Z"
  generateName: tcloud-barad-alarm-amp-7cb448b4c4-
  labels:
    app: tcloud-barad-alarm-amp
    eviction.infra.tce.io/eviction-manager: default
    infra.tce.io/app-name: tcloud-barad-alarm-amp
    infra.tce.io/app-resource-group: "1"
    infra.tce.io/app-type: tad
    infra.tce.io/app-version: 3.10.11-20250214-160805-a57bd60.rhel.amd64
    infra.tce.io/chart-name: tcloud-barad-alarm-amp
    infra.tce.io/comp-name: tcloud-barad-alarm-amp
    infra.tce.io/oam-app: tcloud-barad-alarm-amp
    infra.tce.io/oam-comp: tcloud-barad-alarm-amp
    infra.tce.io/oam-product: barad
    infra.tce.io/oam-project: tce
    infra.tce.io/product: barad
    infra.tce.io/project: tce
    module: tcloud-barad
    pod-template-hash: 7cb448b4c4
    pod.infra.tce.io/pod-spec-hash: 7cb448b4c4
    tcs_app: tcloud-barad-alarm-amp
    tcs_app_type: tad
    tcs_comp: tcloud-barad-alarm-amp
    tcs_instance: tcloud-barad-alarm-amp-7cb448b4c4-qwtmb
    tcs_project: tce
    tcs_region: ap-shenzhen-hqtest-ops
    tcs_resource_group: "1"
    tcs_zone: ap-shenzhen-hqtest-ops-2
  name: tcloud-barad-alarm-amp-7cb448b4c4-qwtmb
  namespace: tce
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: tcloud-barad-alarm-amp-7cb448b4c4
    uid: d7455fef-928f-4844-83d5-a8aef3a1e498
  resourceVersion: "2117756815"
  selfLink: /api/v1/namespaces/tce/pods/tcloud-barad-alarm-amp-7cb448b4c4-qwtmb
  uid: 5a841a5d-136e-4a59-8829-ca939bae517c
spec:
  affinity:
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
      - podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: app
              operator: In
              values:
              - tcloud-barad-alarm-amp
          topologyKey: kubernetes.io/hostname
        weight: 100
  containers:
  - command:
    - /bin/bash
    - -c
    - /usr/bin/python /usr/bin/supervisord -c /etc/supervisord.conf && /usr/sbin/crond
      && tailf /etc/hosts
    env:
    - name: TZ
      value: Asia/Shanghai
    - name: APPLICATION_NAME
      value: tcloud-barad-alarm-amp
    - name: MY_POD_IP
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: status.podIP
    - name: MY_POD_HOSTIP
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: status.hostIP
    - name: MY_POD_NAME
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: metadata.name
    image: registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64
    imagePullPolicy: IfNotPresent
    livenessProbe:
      exec:
        command:
        - bash
        - /tce/healthchk.sh
      failureThreshold: 3
      initialDelaySeconds: 30
      periodSeconds: 20
      successThreshold: 1
      timeoutSeconds: 5
    name: tcloud-barad-alarm-amp
    ports:
    - containerPort: 9411
      name: port-0
      protocol: TCP
    readinessProbe:
      exec:
        command:
        - bash
        - /tce/healthchk.sh
      failureThreshold: 3
      initialDelaySeconds: 30
      periodSeconds: 20
      successThreshold: 1
      timeoutSeconds: 5
    resources:
      limits:
        cpu: "2"
        memory: 4G
      requests:
        cpu: "1"
        memory: 2G
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /tce/conf/global
      name: volume-config
    - mountPath: /tce/conf/cm
      name: volume-local-config
    - mountPath: /data/storage
      name: volume-log
    - mountPath: /tce/customize
      name: volume-customize-confd
      readOnly: true
    - mountPath: /etc/supervisord.d
      name: volume-supervisord-confd
      readOnly: true
    - mountPath: /sys/fs/cgroup
      name: volume-centos-cgroup
      readOnly: true
    - mountPath: /etc/localtime
      name: volume-zoneinfo
      readOnly: true
    - mountPath: /data/customize_packet
      name: volume-customize
    - mountPath: /data/tce.config.center
      name: volume-host-path-0
      readOnly: true
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-t9rmp
      readOnly: true
  dnsConfig:
    options:
    - name: single-request-reopen
    - name: timeout
      value: "1"
  dnsPolicy: ClusterFirst
  enableServiceLinks: false
  nodeName: 10.23.2.180
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: eviction.infra.tce.io/node-offline
    operator: Exists
  topologySpreadConstraints:
  - labelSelector:
      matchLabels:
        app: tcloud-barad-alarm-amp
        module: tcloud-barad
        pod.infra.tce.io/pod-spec-hash: 7cb448b4c4
    maxSkew: 1
    topologyKey: topology.kubernetes.io/zone
    whenUnsatisfiable: DoNotSchedule
  volumes:
  - configMap:
      defaultMode: 420
      items:
      - key: global.json
        mode: 0
        path: global.json
      name: global.conf.d
    name: volume-config
  - configMap:
      defaultMode: 420
      items:
      - key: local.json
        mode: 0
        path: local.json
      name: tcloud-barad-alarm-amp
    name: volume-local-config
  - hostPath:
      path: /data/k8s/log/tce/tcloud-barad-alarm-amp
      type: ""
    name: volume-log
  - configMap:
      defaultMode: 420
      items:
      - key: customize_deploy.sh
        mode: 0
        path: customize_deploy.sh
      - key: customize_move.py
        mode: 0
        path: customize_move.py
      name: global.customize.d
    name: volume-customize-confd
  - configMap:
      defaultMode: 420
      items:
      - key: supervisord.conf
        mode: 0
        path: supervisord.default.ini
      name: global.conf.d
    name: volume-supervisord-confd
  - hostPath:
      path: /sys/fs/cgroup
      type: ""
    name: volume-centos-cgroup
  - hostPath:
      path: /etc/localtime
      type: ""
    name: volume-zoneinfo
  - hostPath:
      path: /data/customize_packet
      type: ""
    name: volume-customize
  - hostPath:
      path: /data/tce.config.center
      type: ""
    name: volume-host-path-0
  - name: default-token-t9rmp
    secret:
      defaultMode: 420
      secretName: default-token-t9rmp
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2025-02-14T09:45:38Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2025-02-14T09:46:15Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2025-02-14T09:46:15Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2025-02-14T09:45:38Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://d7b7c7d3b38944806b4a8d4f52effce7703ca166669c4591c698ea602e5a6d84
    image: registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp:3.10.11-20250214-160805-a57bd60.rhel.amd64
    imageID: docker-pullable://registry.jiguang.woa.com/tcloud-barad-alarm-amp/tcloud-barad-alarm-amp@sha256:dad2955a99dcae40b386e227eacbfcacffee97d262a9d7f56079d775c4314991
    lastState: {}
    name: tcloud-barad-alarm-amp
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2025-02-14T09:45:41Z"
  hostIP: 10.23.2.180
  phase: Running
  podIP: 172.16.13.71
  podIPs:
  - ip: 172.16.13.71
  qosClass: Burstable
  startTime: "2025-02-14T09:45:38Z"
```


获取指定pod的名称的yaml信息
```bash
kubectl get pod pod_name -n tce -o yaml > pod.yaml
```

通过master标签查询环境中的所有master节点信息
```bash
kubectl get nodes --selector=node-role.kubernetes.io/master
```



获取指定名称空间下所有的pod
```bash
kubectl get pods -n namespace_name 

# for example
kubectl get pods -n tce

# 检索指定的pod名称
kubectl get pods -n namespace_name | grep pod_name

# for example
kubectl get pods -n tce | grep policy

```

查看pod启动的详情信息:
```bash
kubectl describe pod pod_name -n namesapce_name

# such as:
kubectl describe pod tcloud-barad-alarm-policy-8467f7d74b-9hcnx -n tce
```

查看pod容器日志
```bash
kubectl logs pod_name -n namespace_name

# such as
kubectl logs tcloud-barad-alarm-policy-8467f7d74b-9hcnx -n tce
```

查看所有的端口信息
```bash
kubectl get entpoints
```

查看所有的namespaces
```bash
kubectl get namespaces
```
