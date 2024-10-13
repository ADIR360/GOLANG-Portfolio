# GoLang-PF
Still Incomplete , will complete it by Novermber
## Description

This project aims to set up an SSH tunnel using ngrok for remote access.I have made a very basic Go-Land Portfolio to test this 


# Curent Position 

The current implementation allows for establishing a connection via SSH, but please note that the SSH server is not functioning properly yet. Further improvements and debugging will be undertaken to ensure a seamless experience.

## Basic Structure

### GoLang Portfolio

- **Note:** This is a very basic structure of code; more **features** will be added by me soon. 
- This repository showcases my GoLang portfolio.  
![Image 1](https://github.com/ADIR360/GoLANG-Portfolio/blob/main/Images/1.png)
- It includes an About Me page.  
![Image 2](https://github.com/ADIR360/GoLANG-Portfolio/blob/main/Images/2.png)
- Highlights my skills in building other cybersecurity applications and tools.  
![Image 3](https://github.com/ADIR360/GoLANG-Portfolio/blob/main/Images/3.png)
- Provides a Contact Me page..  
![Image 5](https://github.com/ADIR360/GoLANG-Portfolio/blob/main/Images/4.png)
- Finally It Provides information about myself.  
![Image 4](https://github.com/ADIR360/GoLANG-Portfolio/blob/main/Images/5.png)


### SSH server

This is a basic structure of the project, and I will be making changes and enhancements as I progress. The following features and functionalities are planned for future updates:

- Improved SSH server configuration
- Enhanced error handling and logging
- User-friendly interface for easier interaction
- Documentation updates and examples for usage

## Getting Started

### Prerequisites

- Ensure you have [ngrok](https://ngrok.com/download) installed on your machine.
- Make sure you have SSH installed and configured on your system.

### Setup

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/yourusername/ssh-tunnel-project.git
   cd ssh-tunnel-project
2. Follow the instructions to configure your SSH server and ngrock 

### Usage 

1. Start SSH server 
2. Run ngrok to create a tunnel
   ```bash
   ngrock tcp 2222
3. Connect SSH tunnel using 
```bash
ssh -p <your_port(2222)> username@localhost

