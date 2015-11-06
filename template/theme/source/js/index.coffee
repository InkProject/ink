require('../css/index.less')
window.jQuery = window.$ = require('jquery')
window.hljs = require('./highlight.pack.js')
require('./jquery.unveil.js')

timeSince = (date) ->

    seconds = Math.floor((new Date() - date) / 1000)
    interval = Math.floor(seconds / 31536000)
    if interval > 1
        return interval + "年前"
    interval = Math.floor(seconds / 2592000)
    if interval > 1
        return interval + "个月前"
    interval = Math.floor(seconds / 86400)
    if interval > 1
        return interval + "天前"
    interval = Math.floor(seconds / 3600)
    if interval > 1
        return interval + "小时前"
    interval = Math.floor(seconds / 60)
    if interval > 1
        return interval + "分钟前"
    return Math.floor(seconds) + "秒前"

$ () ->

    # render date
    $('.date').each (idx, item) ->
        $date = $(item)
        timeStr = $date.data('time')
        if timeStr
            unixTime = Number(timeStr) * 1000
            date = new Date(unixTime)
            $date.prop('title', date).find('.from').text(timeSince(date))
    # render highlight
    $('pre code').each (i, block) ->
        hljs.highlightBlock(block)
    # append image description
    $('img').each (idx, item) ->
        $item = $(item)
        if $item.attr('data-src')
            $item.wrap('<a href="' + $item.attr('data-src') + '" target="_blank"></a>')
            imageAlt = $item.prop('alt')
            $item.parent('a').after('<div class="image-alt">' + imageAlt + '</div>') if $.trim(imageAlt)
    # lazy load images
    if $('img').unveil
        $('img').unveil 200, () ->
            $(@).load () ->
                @style.opacity = 1
