# Helm
You can learn helm from the given Blog and Youtube Videos
- [Blogs](https://blogs.kaiwalyakoparkar.com/what-is-helm)

[![YouTube Video](https://img.youtube.com/vi/fw97GL6LHhg/hqdefault.jpg)](https://youtu.be/fw97GL6LHhg) 




```bash
#Helm chart creation
helm create <name>

#Helm install
helm install <name> <repo-name>

#Listing version
helm list <flag>

#Pulling helm charts
helm pull <chart URL | repo/chartname>

#Pushing helm chart
heml push <name> <repo-name>

# add a chart repository
helm repo add <chart-name> 

#generate an index file given a directory containing packaged charts
helm repo index <chart-name> 

# list chart repositories
helm repo list <chart-name> 

# remove one or more chart repositories
helm repo remove <chart-name> 

# update information of available charts locally from chart repositories
helm repo update <chart-name> 

# search a keywork in chart
helm search repo <keyword>
```



