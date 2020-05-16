
<h1 align="center">Welcome to Termigol 💻</h1>  
  
<div align="center">  
    <img src="images/Termigol.png" alt="The go Gopher"  width=500">  
</div>   
  
  
## Description 📝  
  
I really enjoy learning new stuff, and the easiest path for me is by doing a project.    
I was curious about [Golang](https://golang.org/) and how to use it, so I decided to create a simple terminal emulator.    
> This project does not want in any way be a product, the-best-solution nor compete with any other similar project, this is just a the result of a learning process :stuck_out_tongue:  
  
  
## Goals 🗺️   
  
<details><summary>Commands</summary>
  
<p>

  * [x] ls
  * [x] exit
  * [ ] help
  * [ ] cat
  * [ ] pwd
  * [ ] echo
  * [ ] ping
  * [ ] date
  * [ ] mv
  * [ ] cp

</p>

</details>

<details><summary>Structure</summary>
  
<p>

* [ ]  Have beautiful output colors  
* [ ]  Be able to handle options such `ls -l`  

* Handle commands as just string is not a really good ideia. Maybe we can create a custom `struct` that will handle the command, options, flags, content, etc.  
  
* Being able to perform commands such as `cat file.txt | grep "text"` or `cp file.txt file-bkp.txt && echo "done"`  
  
* Create a graphical interface

</p>

</details>

  

  
## How to use :pencil2:  
  
```go run termigol.go```  
  
## Author  
  
👤 **Vinicius Kammradt**  
  
* Website: https://kammradt.now.sh/  
* Twitter: [@kammzinho](https://twitter.com/kammzinho)  
* Github: [@kammradt](https://github.com/kammradt)  
* LinkedIn: [@vinicius-kammradt](https://linkedin.com/in/vinicius-kammradt)