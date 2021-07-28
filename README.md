# golayout
=======
# A practical Golang project layout

## The tree of file directory
├── bin             //store executable file     
├── cmd             //the number of binaries you want to build.    
├── conf            //project configuration, should be ignored by git.  
├── deployment      //store the Dockerfile  
├── go.mod          
├── Makefile        //the only one make file can build,test,pack,run the project  
├── pkg             //the library used in project  
├── README.md   
├── scripts         //some shell scripts  
└── internal
