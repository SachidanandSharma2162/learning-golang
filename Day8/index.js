express=require("express")
app=express()


app.use(express.json())
app.use(express.urlencoded({extended:true}))
app.get('/',(req,res)=>{
    res.send("Hello")
})
app.post("/signup",(req,res)=>{
    let {userName,password}=req.body
    console.log(userName,password);
    
})
app.listen(3000,()=>{
    console.log("server is running on the port 3000"); 
})