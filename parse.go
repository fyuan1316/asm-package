package main

type ValuesContent struct {
	Global struct {
		Platformscenario string `json:"platformscenario"`
		Replicas         int    `json:"replicas"`
		Cluster          struct {
			Name     string `json:"name"`
			IsGlobal bool   `json:"isGlobal"`
		} `json:"cluster"`
		LabelBaseDomain   string `json:"labelBaseDomain"`
		Namespace         string `json:"namespace"`
		Project           string `json:"project"`
		AdminNotification string `json:"adminNotification"`
		PrometheusName    string `json:"prometheusName"`
		Scheme            string `json:"scheme"`
		Host              string `json:"host"`
		Registry          struct {
			Address          string        `json:"address"`
			ImagePullSecrets []interface{} `json:"imagePullSecrets"`
		} `json:"registry"`
		Ingress struct {
			Host             string `json:"host"`
			IngressClassName string `json:"ingressClassName"`
			Tls              struct {
				SecretName string `json:"secretName"`
			} `json:"tls"`
		} `json:"ingress"`
		ImagesFilter struct {
			Asm     []string `json:"asm"`
			Flagger []string `json:"flagger"`
			Istio   []string `json:"istio"`
		} `json:"images_filter"`
		Images map[string]struct {
				Repository string `json:"repository"`
				Tag        string `json:"tag"`
				Code       string `json:"code"`
				SupportArm bool   `json:"support_arm"`
				Prefix     string `json:"prefix,omitempty"`
		} `json:"images"`
	} `json:"global"`
}