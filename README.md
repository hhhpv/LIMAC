# LIMAC
A Desktop based Video Chat Application

### TECH STACK
* The desktop application is built using [ElectronJS](https://www.electronjs.org/) framework.
* The backend is handled with the help of APIs written in golang and deployed on AWS lambda

#### Important:
* Make sure you download the javascript AWS SDK for your lambda functions and include them in your directory to make the API calls
* Also, Make sure to include your AWS credentials wherever required in the script
* Finally, I have deployed my [TURN server](https://webrtc.org/getting-started/turn-server) on AWS EC2 instance running Ubuntu. You will also need to deploy your TURN sever and 
update the appropriate credentials in your configuration file.

## More About the Application:
* The application uses socket programming to establish peer to peer connection between the clients to make video calls. 
* Here, the application makes use of webRTC protocols to achieve it's goals and establish a real-time connection between two systems with the help of sockets.
* I encourage you to read more about [webRTC here](https://webrtc.org/). 

![LIMAC App](/Screenshots/on_call.jpg)

### Screenshots of the Application:

<details>
<summary>Click here to see screenshots.</summary>
<img src="/Screenshots/start_up.jpg" name="start up"/>
<img src="/Screenshots/login_signup.jpg" name="sign up"/>
<img src="/Screenshots/home_page.jpg" name="home page"/>
<img src="/Screenshots/search_friends.jpg" name="search friends"/>
</details>

## API DOCUMENTATION
### This part contains the API documentation for the go handlers that handle the backend and are deployed on AWS

* Sign Up/login Handler
```
  API Request
  {
    "username":"name_of_user",
    "password":"string"
  }
  
  API Response
  {
    "statusCode":"int",
    "body":{
             "result":"success/failure",
             "message":"string_message"
           }
  }
 ```
 
 * Make Call Handler
```
  API Request
  {
    "caller":"name_of_user",
    "callee":"name_of_user",
    "ip":"callers_ip_address"
  }
  
  API Response
  {
    "statusCode":"int",
    "body":{
             "result":"success/failure",
             "message":"string_message"
           }
  }
 ```
  * List/Find Friend Handler
```
  API Request
  {
    "username":"name_of_user"
  }
  
  API Response on failure
  {
    "statusCode":"int",
    "body":{
             "result":"failure",
             "message":"string_message"
           }
  }
  API Response on success
  {
    "statusCode":"int",
    "body":{
             "result":"failure",
             "message":friendList<StringArray[]>
           }
  }
 ```
  * End Call Handler
```
  API Request
  {
    "callee":"name_of_user",
  }
  
  API Response
  {
    "statusCode":"int",
    "body":{
             "result":"success/failure",
             "message":"string_message"
           }
  }
 ```
   * Check Incoming Call Handler
```
  API Request
  {
    "caller":"name_of_user",
    "callee":"name_of_user"
  }
  
  API Response on failure
  {
    "statusCode":"int",
    "body":{
             "result":"failure",
             "message":"string_message"
           }
  }
  
  API Response on success
  {
    "statusCode":"int",
    "body":{
             "result":"failure",
             "message":"string_message",
             "callerIP":"IP_ADDRESS",
             "caller":"name_of_user"
           }
  }
 ```
 * Add Friend Handler
 ```
    API Request
    {
      "username":"name_of_user",
      "friend":"name_of_user"
    }
    
    API Response
    {
    "statusCode":"int",
    "body":{
             "result":"success/failure",
             "message":"string_message"
           }
    }
 ```
 
 ## Important:
 **There is more to the application apart from the documentation check out the desktop app folder and go through the code to understand how sockets are implemented and how the webRTC
 protocols are implemented. You can also read more about how to deploy a simple go handler on AWS lambda on [my blog here](https://hitheshpv.net/2021/01/16/deploying-a-simple-go-handler-aws-lambda/).**
