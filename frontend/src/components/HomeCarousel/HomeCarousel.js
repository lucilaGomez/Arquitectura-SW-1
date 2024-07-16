import React, { useEffect } from 'react';
import Splide from '@splidejs/splide';
import '@splidejs/splide/dist/css/splide.min.css';
import './HomeCarousel.css';

const HomeCarousel = ({ images }) => {
  useEffect(() => {
    var main = new Splide('#main-slider', {
      type: 'fade',
      heightRatio: 0.5,
      pagination: false,
      arrows: false,
      cover: true,
    });

    var thumbnails = new Splide('#thumbnail-slider', {
      rewind: true,
      fixedWidth: 104,
      fixedHeight: 58,
      isNavigation: true,
      gap: 10,
      focus: 'center',
      pagination: false,
      cover: true,
      dragMinThreshold: {
        mouse: 4,
        touch: 10,
      },
      breakpoints: {
        640: {
          fixedWidth: 66,
          fixedHeight: 38,
        },
      },
    });

    main.sync(thumbnails);
    main.mount();
    thumbnails.mount();
  }, [images]);

  return (
    <div className="carousel-wrapper">
      <div id="main-slider" className="splide">
        <div className="splide__track">
          <ul className="splide__list">
            {images.map((url, index) => (
              <li className="splide__slide" key={index}>
                <img src={url} alt={`Slide ${index}`} />
              </li>
            ))}
          </ul>
        </div>
      </div>

      <div id="thumbnail-slider" className="splide">
        <div className="splide__track">
          <ul className="splide__list">
            {images.map((url, index) => (
              <li className="splide__slide" key={index}>
                <img src={url} alt={`Thumbnail ${index}`} />
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default HomeCarousel;
