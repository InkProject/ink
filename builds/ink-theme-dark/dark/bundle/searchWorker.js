function load(url, callback) {
	var xhr
	if (typeof XMLHttpRequest !== 'undefined') xhr = new XMLHttpRequest()
	else {
		var versions = ["MSXML2.XmlHttp.5.0",
		 	"MSXML2.XmlHttp.4.0",
		 	"MSXML2.XmlHttp.3.0",
		 	"MSXML2.XmlHttp.2.0",
		 	"Microsoft.XmlHttp"]

		for(var i = 0, len = versions.length; i < len; i++) {
  		try {
  			xhr = new ActiveXObject(versions[i])
  			break
  		} catch(e) {}
		}
	}
	xhr.onreadystatechange = function() {
		if(xhr.readyState < 4) return
		if(xhr.status !== 200) return
		if(xhr.readyState === 4) callback(xhr)
	}
	xhr.open('GET', url, true)
	xhr.send('')
}

var data, lastKeyword

var search = function(keyword) {
  var results = []
  var keywordStr = keyword.toLowerCase().replace(/ /g, ' ')
  var keywords = keywordStr.split(' ')
  for (var i = 0; i < data.length; i++) {
    var item = data[i]
    var title = item.title.toLowerCase()
    var preview = item.preview.toLowerCase()
    var content = item.content.toLowerCase()
    var isMatch = true
    for (var j = 0; j < keywords.length; j++) {
      var key = keywords[j]
      if (title.indexOf(key) == -1 &&
        preview.indexOf(key) == -1 &&
        content.indexOf(key) == -1) {
        isMatch = false
        break
      }
    }
    if (isMatch) results.push(item)
  }
  postMessage({
    keyword: keywordStr,
    keywords: keywords,
    results: results
  })
}


onmessage = function(event) {
  var action = event.data.action
  if (action == 'start') {
    load(event.data.root + '/index.json', function(xhr) {
      data = JSON.parse(xhr.responseText)
      if (lastKeyword) {
        search(lastKeyword)
      }
    })
  } else {
    if (data) {
      search(event.data.keyword)
    } else {
      lastKeyword = event.data.keyword
    }
  }
}
