import express from "express"

const app = express()

app.get("/", (req, res) => {
  res.json({ msg: "hello from express server" }).status(200);
})

app.listen(3001, () => {
  console.log("Express server is running on PORT: 3001")
})