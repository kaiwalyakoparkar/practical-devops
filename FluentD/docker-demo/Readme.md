# FluentD

- So basically it first starts a container with the `app.sh` application
- This application sends logs [Saves to file when file], as [Http with http request]
- These logs are detected by fleuntd based on their types we describe in configurate
- And hense stores them in different location outputs as we mention in Configuration