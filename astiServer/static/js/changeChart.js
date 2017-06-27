var slide = document.querySelector("#slide");
slide.addEventListener("input", function(){
  var pie = document.querySelector("#pie");
  var newChart = "static/pieCharts/pieChart"+ slide.value +".png";
  pie.setAttribute("src", newChart);
});
