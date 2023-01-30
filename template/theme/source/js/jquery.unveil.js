/**
 * jQuery Unveil
 * A very lightweight jQuery plugin to lazy load images
 * http://luis-almeida.github.com/unveil
 *
 * Licensed under the MIT license.
 * Copyright 2013 Luís Almeida
 * https://github.com/luis-almeida
 * 
 * 
 * Modified by @w568w to work with the lastest jQuery.
 */
import $ from 'jquery'


(function () {
  $.fn.unveil = function (threshold, callback) {

    let $w = $(window),
      th = threshold || 0,
      retina = window.devicePixelRatio > 1,
      attrib = retina ? "data-src-retina" : "data-src",
      images = this,
      loaded;

    this.one("unveil", function () {
      let source = this.getAttribute(attrib);
      source = source || this.getAttribute("data-src");
      if (source) {
        this.setAttribute("src", source);
        if (typeof callback === "function") callback.call(this);
      }
    });

    function unveil() {
      const inview = images.filter(function () {
        let $e = $(this);
        if ($e.is(":hidden")) return;

        let wt = $w.scrollTop(),
          wb = wt + $w.height(),
          et = $e.offset().top,
          eb = et + $e.height();

        return eb >= wt - th && et <= wb + th;
      });

      loaded = inview.trigger("unveil");
      images = images.not(loaded);
    }

    $w.on("scroll.unveil resize.unveil lookup.unveil", unveil);

    unveil();

    return this;

  };
})();