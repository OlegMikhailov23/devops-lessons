{{- define "momo-backend.name" -}}
{{- .Chart.Name -}}
{{- end -}}

{{- define "momo-backend.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "momo-backend.labels" -}}
app.kubernetes.io/name: {{ include "momo-backend.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | quote }}
app.kubernetes.io/component: backend
app.kubernetes.io/part-of: momo-store
env: {{ .Values.env.name | default "production" }}
{{- end }}
