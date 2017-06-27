{{define "index"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="static/css/bootstrap.min.css" />
  </head>
  <body>
    <div class="container">
      <div class="row">
        <div class="col-sm-10 col-sm-offset-1">
          <h1 class="panel-title text-center">Welcome to my worm wars simulator!</h1>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-10 col-sm-offset-1">
          <form action="/sim" method="post">
            <div class="form-group">
            <label>Population Size:</label>
            <input type="number" name="pop" class="form-control" value=0><br />
            </div>
            <div class="form-group">
            <label>Initial Seed:</label>
            <input type="number" name="initI" class="form-control" value=0><br />
            </div>
            <div class="row">
              <div class="col-sm-4">
                <div class="form-group">
                  <label>S -> I:</label>
                  <input type="number" name="S2I" class="form-control" placeholder="0.1" step="0.01" min="0" max="1"/><br />
                </div>
              </div>
              <div class="col-sm-4">
                <div class="form-group">
                  <label>I -> R:</label>
                  <input type="number" name="I2R" class="form-control" placeholder="0.1" step="0.01" min="0" max="1"/><br />
                </div>
              </div>
              <div class="col-sm-4">
                <div class="form-group">
                  <label>S -> R:</label>
                  <input type="number" name="S2R" class="form-control" required placeholder="0.1" step="0.01" min="0" max="1"/><br />
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-sm-8 col-sm-offset-4">
              <input type="submit" value="Run the simulation!" class="btn btn-default">
            </div>
          </div>
          </form>
        </div>
      </div>
    </div>
  </body>
</html>
{{end}}
