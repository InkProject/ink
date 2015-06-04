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
        $item.attr('src', 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAB+UlEQVRYR+2XQSgFQRjH30sppUgp5SDlJqLkLMVJSinF5TkppaSUg5QiclDKhYgLcSEHOUgpUopycXBSclbipMT/r9nXmHa93ZnvrUc79W93Z+f7vt/+Z2fe23SqwFu6wPlSCaDrDCUOJg66OuAa/2ffwQo8+QrU6+pAxPhtjB+Bnry4IAdPMKA9YnKp4YdI1JUL8EOqmkWeZ8SURwGM6z31TPnfgEWwfgzqh96gVWjdYhr9QkQcnEbmKSP7IK43Q0BmMIbTtx8wVgTwEcmrjQLXuG7JAdiE+xcKsB7H7DaixYkAviBhqQFzg+vmHwDLcO8KqlNj6DZdN5sI4AKyjhuZubEua33dOG+AZlXfAY7s01snLo6NPhHAEiSdgfgLw0XCBTKvFWrF+SnEccMQ3ZvzcesefY3Qq3ZPBNCnVrarFmeXUKU26B3nXPl+bQmdo3EB0qkziFMbthG+DTpXAXlzsBgFjiCb3+47NdV8XcQAO5RTi+rJt3Dkxm3buJgmpQA5lbcQ90IC8sknbMlU3AOONVKAG0iUcQQyw8UAuZdxT5NuIoD8h82prZKmQz4RwB0k6ssDHFOKAOaJ7Vta620mDji9Rqh/1HuI6ImbTNXjl92AVzvoe4OLYu0XIHdRcwiii18trg8i68lIAK2tS6bY1brEQSEHPwHGT3gpenNNIQAAAABJRU5ErkJggg==')
        $item.wrap('<a href="' + $item.attr('data-src') + '" target="_blank"></a>')
        imageAlt = $item.prop('alt')
        $item.parent('a').after('<div class="image-alt">' + imageAlt + '</div>') if $.trim(imageAlt)
    # lazy load images
    .unveil 200, () ->
        $(@).load () ->
            @style.opacity = 1
