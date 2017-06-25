{{define "index"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <style>
    html, body {
    	margin: 0 auto;
    	padding: 0 auto;
    	font-family: "Times New Roman", arial, Verdana;
    	color: #817a53;
    	background: #d8d6cb;
    }

    label {
      display: inline-block;
      width: 140px;
      text-align: right;
    }​
    </style>
  </head>
  <body>
    <div style="text-align:center;">
      <p>Welcome to my worm wars simulator!</p>
      <form action="/sim" method="post">
        <label>Population Size:</label><input type="text" name="pop"><br />
  			<label>Initial Seed:</label><input type="text" name="initI"><br />
        <label>S -> I:</label><input type="text" name="S2I" /><br />
        <label>I -> R:</label><input type="text" name="I2R"><br />
        <label>S -> R:</label><input type="text" name="S2R"><br />
  			<input type="submit" value="Run the simulation!">
      </form>
    </div>
  </body>
</html>
{{end}}
