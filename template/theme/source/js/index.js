require('../css/index.css')
window.jQuery = window.$ = require('jquery')
window.hljs = require('./highlight.pack.js')
require('./jquery.unveil.js')

var timeSince = function(date) {
  var seconds = Math.floor((new Date() - date) / 1000)
  var interval = Math.floor(seconds / 31536000)
  if (interval > 1) return interval + timeSinceLang.year

  interval = Math.floor(seconds / 2592000)
  if (interval > 1) return interval + timeSinceLang.month

  interval = Math.floor(seconds / 86400)
  if (interval > 1) return interval + timeSinceLang.day

  interval = Math.floor(seconds / 3600)
  if (interval > 1) return interval + timeSinceLang.hour

  interval = Math.floor(seconds / 60)
  if (interval > 1) return interval + timeSinceLang.minute

  return Math.floor(seconds) + timeSinceLang.second
}

$(function() {
  // render date
  $('.date').each(function(idx, item) {
    var $date = $(item)
    var timeStr = $date.data('time')
    if (timeStr) {
      var unixTime = Number(timeStr) * 1000
      var date = new Date(unixTime)
      $date.prop('title', date).find('.from').text(timeSince(date))
    }
  })
  // render highlight
  $('pre code').each(function(i, block) {
    hljs.highlightBlock(block)
  })
  // append image description
  $('img').each(function(idx, item) {
    $item = $(item)
    if ($item.attr('data-src')) {
      $item.wrap('<a href="' + $item.attr('data-src') + '" target="_blank"></a>')
      var imageAlt = $item.prop('alt')
      if ($.trim(imageAlt)) $item.parent('a').after('<div class="image-alt">' + imageAlt + '</div>')
    }
  })
  // lazy load images
  if ($('img').unveil) {
    $('img').unveil(200, function() {
      $(this).load(function() {
        this.style.opacity = 1
      })
    })
  }
})
