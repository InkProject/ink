(function() {
  var timeSince;

  timeSince = function(date) {
    var interval, seconds;
    seconds = Math.floor((new Date() - date) / 1000);
    interval = Math.floor(seconds / 31536000);
    if (interval > 1) {
      return interval + " years ago";
    }
    interval = Math.floor(seconds / 2592000);
    if (interval > 1) {
      return interval + " months ago";
    }
    interval = Math.floor(seconds / 86400);
    if (interval > 1) {
      return interval + " days ago";
    }
    interval = Math.floor(seconds / 3600);
    if (interval > 1) {
      return interval + " hours ago";
    }
    interval = Math.floor(seconds / 60);
    if (interval > 1) {
      return interval + " mins ago";
    }
    return Math.floor(seconds) + " seconds ago";
  };

  $(function() {
    $('.date').each(function(idx, item) {
      var $date, date, timeStr, unixTime;
      $date = $(item);
      timeStr = $date.data('time');
      if (timeStr) {
        unixTime = Number(timeStr) * 1000;
        date = new Date(unixTime);
        return $date.prop('title', date).find('.from').text(timeSince(date));
      }
    });
    $('pre code').each(function(i, block) {
      return hljs.highlightBlock(block);
    });
    $('img').each(function(idx, item) {
      var $item, imageAlt;
      $item = $(item);
      if ($item.attr('data-src')) {
        $item.wrap('<a href="' + $item.attr('data-src') + '" target="_blank"></a>');
        imageAlt = $item.prop('alt');
        if ($.trim(imageAlt)) {
          return $item.parent('a').after('<div class="image-alt">' + imageAlt + '</div>');
        }
      }
    });
    if ($('img').unveil) {
      return $('img').unveil(200, function() {
        return $(this).load(function() {
          return this.style.opacity = 1;
        });
      });
    }
  });

}).call(this);
