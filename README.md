# idor_enum

Golang script to automate the IDOR enumeration process.

To use this script, just have any golang version installed.

This script now makes GET and POST requests for each ID and prints the response. 

Error handling has been improved to handle errors when creating the request, making the request, and reading the response. 

It also checks the response status code and prints an error if it is not 200 (OK).

<div align="center">
  <br/>
  <img src="https://github.com/washingtonP1974/idor_enum/blob/main/1.png" alt="1">
</div>


This version of the script that does not require authentication
 
##   Standard:

Example:

go run main.go http://ip target host:port

go run main.go http://54.172.238.xxx:80 



Remember, it is important to use IDOR enumeration tools responsibly and ethically. 

Always obtain proper permission before testing any systems that you do not own or control.

Feel free to optimize and express your suggestion, thank you.
