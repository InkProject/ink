import '../css/index.css'
import $ from 'jquery'
import './jquery.unveil'
import searchTpl from './searchTpl.html'

// backward compatibility for old version of highlight.js
window.hljs = require('./highlight.pack.js')

// pick from underscore
let debounce = function (func, wait, immediate) {
  let timeout, args, context, timestamp, result;

  let later = function () {
    let last = Date.now() - timestamp;

    if (last < wait && last >= 0) {
      timeout = setTimeout(later, wait - last);
    } else {
      timeout = null;
      if (!immediate) {
        result = func.apply(context, args);
        if (!timeout) context = args = null;
      }
    }
  };

  return function () {
    context = this;
    args = arguments;
    timestamp = Date.now();
    let callNow = immediate && !timeout;
    if (!timeout) timeout = setTimeout(later, wait);
    if (callNow) {
      result = func.apply(context, args);
      context = args = null;
    }

    return result;
  };
}

let timeSince = function (date) {
  let seconds = Math.floor((new Date() - date) / 1000)
  let interval = Math.floor(seconds / 31536000)
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

let initSearch = function () {
  let searchDom = $('#search')
  if (!searchDom.length) return
  let searchWorker = new Worker(root + '/bundle/searchWorker.js')
  let oriHtml = $('.article-list').html()
  let workerStarted = false
  let tpl = function (keywords, title, preview, link, cover) {
    for (let i = 0; i < keywords.length; i++) {
      let keyword = keywords[i]
      let wrap = '<span class="searched">' + keyword + '</span>'
      let reg = new RegExp(keyword, 'ig')
      title = title.replace(reg, wrap)
      preview = preview.replace(reg, wrap)
    }
    return searchTpl
      .replace('{{title}}', title)
      .replace('{{link}}', link + '?search=' + keywords)  // append keywords to url
      .replace('{{preview}}', preview)
  }
  searchWorker.onmessage = function (event) {
    let results = event.data.results
    let keywords = event.data.keywords
    if (results.length) {
      let retHtml = ''
      for (let i = 0; i < results.length; i++) {
        let item = results[i]
        let itemHtml = tpl(keywords, item.title, item.preview, item.link, item.cover)
        retHtml += itemHtml
      }
      $('.page-nav').hide()
      $('.article-list').html(retHtml)
    } else {
      let keyword = event.data.keyword
      if (keyword) {
        $('.page-nav').hide()
        $('.article-list').html('<div class="empty">未搜索到 "<span>' + keyword + '</span>"</div>')
      } else {
        $('.page-nav').show()
        $('.article-list').html(oriHtml)
      }
    }
  }
  searchDom.on('input', debounce(function () {
    let keyword = $(this).val().trim()
    if (keyword) {
      searchWorker.postMessage({
        search: 'search',
        keyword: keyword
      })
    } else {
      $('.page-nav').show()
      $('.article-list').html(oriHtml)
    }
  }, 500))
  searchDom.on('focus', function () {
    if (!workerStarted) {
      searchWorker.postMessage({
        action: 'start',
        root: root
      })
      workerStarted = true
    }
  })
}

$(function () {
  // render date
  $('.date').each(function (idx, item) {
    let $date = $(item)
    let timeStr = $date.data('time')
    if (timeStr) {
      let unixTime = Number(timeStr) * 1000
      let date = new Date(unixTime)
      $date.prop('title', date).text(timeSince(date))
    }
  })
  // render highlight
  $('pre code').each(function (_, block) {
    hljs.highlightBlock(block)
  })
  // append image description
  $('img').each(function (_, item) {
    let $item = $(item)
    if ($item.attr('data-src')) {
      $item.wrap('<a href="' + $item.attr('data-src') + '" target="_blank"></a>')
      let imageAlt = $item.prop('alt')
      if (imageAlt.trim()) $item.parent('a').after('<div class="image-alt">' + imageAlt + '</div>')
    }
  })
  // lazy load images
  if ($('img').unveil) {
    $('img').unveil(200, function () {
      $(this).on("load", function () {
        this.style.opacity = 1
      })
    })
  }
  // init search
  initSearch()
})

/**
 * Get all text nodes under an element
 * @param {!Element} el
 * @return {Array<!Node>}
 */
function getTextNodes(el) {
  const iterator = document.createNodeIterator(el, NodeFilter.SHOW_TEXT);
  const textNodes = [];
  let currentTextNode;
  while ((currentTextNode = iterator.nextNode())) {
    textNodes.push(currentTextNode);
  }
  return textNodes;
}

// get searched keywords from url
let reg = new RegExp("(^|&)search=([^&]*)(&|$)")
let r = window.location.search.substring(1).match(reg)
if (r != null && r.toString().length > 1) {
  let keywords = decodeURI(r[2]).split(',')
  // highlight searched keywords
  getTextNodes(document.body).forEach(function (node) {
    let i;
    for (i = 0; i < node.parentNode.childNodes.length; i++) {
      if (node.parentNode.childNodes[i] === node) {
        break;
      }
    }

    let content = node.nodeValue
    let oldContent = content
    for (let i = 0; i < keywords.length; i++) {
      let keyword = keywords[i]
      let wrap = '<span class="searched">' + keyword + '</span>'
      let reg = new RegExp(keyword, 'ig')
      content = content.replace(reg, wrap)
    }

    // replace the text node with a span containing the highlighted text
    if (content != oldContent) {
      let newNode = document.createElement('span')
      newNode.innerHTML = content
      node.parentNode.replaceChild(newNode, node)
    }
  })
  // scroll to the first searched keyword
  let elmnt = document.getElementsByClassName("searched")
  elmnt[0].scrollIntoView()
}

// show "go top" button when needed
$(function () {
  $("#go_top").hide();
  $(function () {
    let height = $(window).height();
    $(window).on("scroll", function () {
      if ($(window).scrollTop() > height) {
        $("#go_top").fadeIn(500);
      } else {
        $("#go_top").fadeOut(500);
      }
    });
    $("#go_top").on("click", function () {
      $('body,html').animate({ scrollTop: 0 }, 100);
      return false;
    });
  });
});