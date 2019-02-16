var election = new Datamap({
    scope: 'usa',
    element: document.getElementById('container'),
    responsive: true,
    geographyConfig: {
      highlightBorderColor: '#ddd',
     popupTemplate: function(geography, data) {
        return '<div class="hoverinfo">' + geography.properties.name + ' contacted: ' +  data.hits + ' times'
      },
      highlightBorderWidth: 2,
      highlightFillColor: '#FF69B4'
    },
  
    fills: {
        'min':'#00acd3',
        'mid':'#00849b',
        'max':'#444749',
    defaultFill: '#7bc6eb',
    
  },
  data:{
    "AZ": {
        "fillKey": "min",
        "hits": 5
    },
    "CO": {
        "fillKey": "mid",
        "hits": 5
    },
    "DE": {
        "fillKey": "mid",
        "hits": 32
    },
    "FL": {
        "fillKey": "UNDECIDED",
        "hits": 29
    },
    "GA": {
        "fillKey": "min",
        "hits": 32
    },
    "HI": {
        "fillKey": "mid",
        "hits": 32
    },
    "ID": {
        "fillKey": "min",
        "hits": 32
    },
    "IL": {
        "fillKey": "mid",
        "hits": 32
    },
    "IN": {
        "fillKey": "min",
        "hits": 11
    },
    "IA": {
        "fillKey": "mid",
        "hits": 11
    },
    "KS": {
        "fillKey": "min",
        "hits": 32
    },
    "KY": {
        "fillKey": "min",
        "hits": 32
    },
    "LA": {
        "fillKey": "min",
        "hits": 32
    },
    "MD": {
        "fillKey": "mid",
        "hits": 32
    },
    "ME": {
        "fillKey": "mid",
        "hits": 32
    },
    "MA": {
        "fillKey": "mid",
        "hits": 32
    },
    "MN": {
        "fillKey": "mid",
        "hits": 32
    },
    "MI": {
        "fillKey": "mid",
        "hits": 32
    },
    "MS": {
        "fillKey": "min",
        "hits": 32
    },
    "MO": {
        "fillKey": "min",
        "hits": 13
    },
    "MT": {
        "fillKey": "min",
        "hits": 32
    },
    "NC": {
        "fillKey": "max",
        "hits": 32
    },
    "NE": {
        "fillKey": "min",
        "hits": 32
    },
    "NV": {
        "fillKey": "max",
        "hits": 32
    },
    "NH": {
        "fillKey": "mid",
        "hits": 32
    },
    "NJ": {
        "fillKey": "mid",
        "hits": 32
    },
    "NY": {
        "fillKey": "mid",
        "hits": 32
    },
    "ND": {
        "fillKey": "min",
        "hits": 32
    },
    "NM": {
        "fillKey": "mid",
        "hits": 32
    },
    "OH": {
        "fillKey": "mid",
        "hits": 32
    },
    "OK": {
        "fillKey": "min",
        "hits": 32
    },
    "OR": {
        "fillKey": "mid",
        "hits": 32
    },
    "PA": {
        "fillKey": "mid",
        "hits": 32
    },
    "RI": {
        "fillKey": "mid",
        "hits": 32
    },
    "SC": {
        "fillKey": "min",
        "hits": 32
    },
    "SD": {
        "fillKey": "min",
        "hits": 32
    },
    "TN": {
        "fillKey": "min",
        "hits": 32
    },
    "TX": {
        "fillKey": "min",
        "hits": 32
    },
    "UT": {
        "fillKey": "min",
        "hits": 32
    },
    "WI": {
        "fillKey": "mid",
        "hits": 32
    },
    "VA": {
        "fillKey": "mid",
        "hits": 32
    },
    "VT": {
        "fillKey": "mid",
        "hits": 32
    },
    "WA": {
        "fillKey": "mid",
        "hits": 32
    },
    "WV": {
        "fillKey": "min",
        "hits": 32
    },
    "WY": {
        "fillKey": "min",
        "hits": 32
    },
    "CA": {
        "fillKey": "mid",
        "hits": 32
    },
    "CT": {
        "fillKey": "mid",
        "hits": 32
    },
    "AK": {
        "fillKey": "min",
        "hits": 32
    },
    "AR": {
        "fillKey": "min",
        "hits": 32
    },
    "AL": {
        "fillKey": "min",
        "hits": 32
    }
  }
  });
  election.labels();