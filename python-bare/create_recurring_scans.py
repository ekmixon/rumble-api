import requests
import os
import json

token = os.getenv("RUMBLE_API_KEY", "")
if token == "":
    print("Error: The RUMBLE_API_KEY environment variable is not set")
    exit(1)

baseURL = "https://console.rumble.run/api/v1.0"
headers = {
    "Authorization": f"Bearer {token}",
    "Content-Type": "application/json",
}


# Make sure our key is valid
keyInfo = requests.get(f"{baseURL}/org/key", data={}, headers=headers).json()
print("API Key: ", keyInfo)

if 'organization_id' not in keyInfo:
    print("Error: The RUMBLE_API_KEY value is not a valid organization API key")
    exit(1)


agentList = requests.get(
    f"{baseURL}/org/agents", data={}, headers=headers
).json()

print("Found %d agents" % (len(agentList)))

taskList = requests.get(
    f"{baseURL}/org/tasks?status=active", data={}, headers=headers
).json()

print("Found %d active tasks" % (len(taskList)))

recurringTaskList = [task for task in taskList if task["recur"]]
print("Found %d recurring tasks" % (len(recurringTaskList)))

siteList = requests.get(
    f"{baseURL}/org/sites", data={}, headers=headers
).json()

print("Found %d sites" % (len(siteList)))


# Use the description to determine if a recurring scan has already been scheduled for this site
for site in siteList:
    if site["description"] and "api-managed" not in site["description"]:
        # Create a recurring task for this site
        # TODO: Make sure we pick the right agent (local mapping or otherwise)
        scan = requests.put(
            f"{baseURL}/org/sites/" + site["id"] + "/scan",
            json={
                "targets": site["scope"],
                "rate": "6000",
                "scan-frequency": "daily",
            },
            headers=headers,
        ).json()

        print("Scan created for site " + site["name"] + f": {scan}")

        # Update the site to indicate we have a task
        update = requests.patch(
            f"{baseURL}/org/sites/" + site["id"],
            json={"description": "api-managed site"},
            headers=headers,
        ).json()

        print("Updated site " + site["name"] + f": {update}")

