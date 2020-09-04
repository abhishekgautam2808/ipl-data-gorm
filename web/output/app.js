
fetch('json1.json')
.then((response) => response.json())
.then((response) => FirstBarChart(response));

fetch('json2.json')
.then((response) => response.json())
.then((response) => SecondBarChart(response));

fetch('json3.json')
.then((response) => response.json())
.then((response) => ThirdBarChart(response));


fetch('json4.json')
.then((response) => response.json())
.then((response) => FourthBarChart(response));



function FirstBarChart(data)
{  //The Object.entries() method returns an array of a given object
    const checkData = Object.entries(data).map(i=>(  
         {
            name:i[0],
            y:i[1],
        }
    ));
  
Highcharts.chart('totalRunsByTeam', {
    chart: {
        type: 'column'
    },
    title: {
        text: 'Total Run\'s By Team'
    },
    subtitle: {
        text: ''
    },
    xAxis: {
        type: 'category',
        labels: {
            rotation: +90,
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    },
    yAxis: {
        min: 0,
        title: {
            text: 'Total Runs\'s (Thousands)'
        }
    },
    legend: {
        enabled: false
    },
    tooltip: {
        pointFormat: 'Total Run\'s: <b>{point.y:.0f} Thousand</b>'
    },
    series: [{
        name: 'Run\'s',
         data: checkData,
        dataLabels: {
            enabled: true,
            rotation: -90,
            color: '#FFFFFF',
            align: 'right',
            format: '{point.y:.1f}', // one decimal
            y: 10, // 10 pixels down from the top
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    }]
});

}

function SecondBarChart(data)
{
    const checkData = Object.entries(data).map(i=>{
        return {
            name:i[0],
            y:i[1],
        };
    });
    
Highcharts.chart('topBatsman', {
    chart: {
        type: 'column'
    },
    title: {
        text: 'Total Run\'s By RCB player\'s'
    },
    subtitle: {
        text: ''
    },
    xAxis: {
        type: 'category',
        labels: {
            rotation: +90,
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    },
    yAxis: {
        min: 0,
        title: {
            text: 'Total Runs\'s (Thousands)'
        }
    },
    legend: {
        enabled: false
    },
    tooltip: {
        pointFormat: 'Total Run\'s: <b>{point.y:.0f}</b>'
    },
    series: [{
        name: 'Run\'s',
         data: checkData,
        dataLabels: {
            enabled: true,
            rotation: -90,
            color: '#FFFFFF',
            align: 'right',
            format: '{point.y:.1f}', // one decimal
            y: 10, // 10 pixels down from the top
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    }]
});

}




function ThirdBarChart(data)
{
    const checkData = Object.entries(data).map(i=>{
        return {
            name:i[0],
            y:i[1],
        };
    });
    
Highcharts.chart('foreignUmpire', {
    chart: {
        type: 'column'
    },
    title: {
        text: 'Total Matches By Foreign Umpires'
    },
    subtitle: {
        text: ''
    },
    xAxis: {
        type: 'category',
        labels: {
            rotation: +0,
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    },
    yAxis: {
        min: 0,
        title: {
            text: 'Total Matches '
        }
    },
    legend: {
        enabled: false
    },
    tooltip: {
        pointFormat: 'Total Run\'s: <b>{point.y:.0f}</b>'
    },
    series: [{
        name: 'Run\'s',
         data: checkData,
        dataLabels: {
            enabled: true,
            rotation: -90,
            color: '#FFFFFF',
            align: 'right',
            format: '{point.y:.1f}', // one decimal
            y: 10, // 10 pixels down from the top
            style: {
                fontSize: '13px',
                fontFamily: 'Verdana, sans-serif'
            }
        }
    }]
});

}


function FourthBarChart(response){


var years= ["2008","2009","2010","2011","2012","2013","2014","2015","2016","2017"]  
var keys = Object.keys(response)
var temp= Object.entries(response).map(i=>(
    { 
        name : i[0],
        data : Object.values(i[1]),

    }
));

Highcharts.chart('matchesByTeam', {
    chart: {
        type: 'column'
    },
    title: {
        text: 'Matches by team in different seasons'
    },
    subtitle: {
        text: ''
    },
    xAxis: {
        
        categories: years,
        crosshair: true
    },
    yAxis: {
        min: 0,
        title: {
            text: 'Matches'
        }
    },
    
    plotOptions: {
        column: {
            pointPadding: 0.2,
            borderWidth: 0
        }
    },
    
    series: temp
});
              
}