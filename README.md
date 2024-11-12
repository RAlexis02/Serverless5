# Serverless5

Serverless5 Project - Go Web Application
This project, Serverless5, is a simple web application built with Go that displays "Hola Mundo". It has been deployed on Railway for easy access and demonstrates basic Go deployment in a serverless environment.

📋 Project Overview
Serverless5 is a basic Go web application that responds with "Hola Mundo". This project is ideal for beginners who want to learn how to set up and deploy a Go application in a serverless environment using Railway.

🛠 Requirements
To deploy this project on Railway, you’ll need:

A GitHub account with the project files stored in a repository.
A Railway account. You can sign up at Railway.app.
The main project file: main.go.
📂 Project Structure
plaintext
Copiar código
Serverless5/
├── main.go           # Main Go file that displays "Hola Mundo"
└── README.md         # Documentation file with project information and deployment guide
main.go: Contains the Go code to display "Hola Mundo" when accessed via a web browser.
README.md: Provides detailed instructions on how to deploy the project on Railway.
🚀 Step-by-Step Deployment on Railway
Follow these steps to deploy Serverless5 on Railway.

Step 1: Set Up the Project in GitHub
Create a GitHub Repository: If you haven’t already, create a repository on GitHub for Serverless5.
Upload Project Files: Ensure that main.go is in the root directory of the repository.
Commit Changes: Commit and push this file to the main branch on GitHub.
Step 2: Set Up a Railway Account and Connect GitHub
Sign Up/In to Railway: Go to Railway.app and log in or create an account.
Connect GitHub: In Railway’s dashboard, link your GitHub account. This will allow Railway to access your repositories for deployment.
Step 3: Deploy the Project on Railway
Create a New Project:

On the Railway dashboard, click on New Project.
Select Deploy from GitHub repo to deploy from an existing GitHub repository.
Select the Repository:

Choose the Serverless5 repository from the list.
Railway will detect that it’s a Go project and will prepare the environment accordingly.
Configuring the Port to 8080:

Go to the Settings tab in Railway and scroll to the Networking section.
Under Public Networking, click Generate Domain.
If prompted, enter 8080 as the target port for the application.
Railway will generate a public URL for accessing your application.
Step 4: Access the Application
Open the Generated URL:
Railway will provide a URL like https://serverless5-production.up.railway.app. Copy and open this URL in a browser to view the application.
Verify the Functionality:
You should see the message "Hola Mundo desde Go en Railway!" displayed on the page.
🔄 Summary of Commands for GitHub
For quick reference, here are the basic Git commands to upload your project to GitHub.

bash
Copiar código
# Initialize a new Git repository (if needed)
git init

# Add all project files
git add .

# Commit the changes
git commit -m "Initial commit of Serverless5 project"

# Push to GitHub (replace 'main' if your branch name is different)
git push origin main
💡 Troubleshooting Tips
If you encounter issues while deploying, here are some things to check:

Root Directory: Ensure that main.go is in the root directory of your repository.
Public Networking: Verify that you generated a domain under Public Networking in Railway’s Settings.
Port Configuration: Make sure to specify port 8080 if prompted when generating the domain in Railway.
📜 License
This project is open-source and available under the MIT License.