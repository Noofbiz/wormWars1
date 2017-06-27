{{define "sim"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="static/css/bootstrap.min.css" />
  </head>
  <body>
    <div class="container">
      <div class="row">
        <div class="col-sm-10 col-sm-offset-1">
          <h1 class="text-center"> Simulation Results! </h1>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-10 col-sm-offset-1">
          <img id="pie" src="static/pieCharts/pieChart0.png" class="img-responsive center-block">
        </div>
      </div>
      <div class="row">
        <div class="col-sm-10 col-sm-offset-1">
          <input id="slide" type="range" value="0" min="0" max="{{.Max}}" step="1" />
        </div>
      </div>
    </div>
    <script src="static/js/changeChart.js"></script>
  </body>
</html>
{{end}}
