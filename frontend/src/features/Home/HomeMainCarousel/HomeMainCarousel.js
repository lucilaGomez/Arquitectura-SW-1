import HomeCarousel from "../../../components/HomeCarousel/HomeCarousel";

const url = [
  "https://www.cadenadial.com/wp-content/uploads/2021/10/saad-khan-425b2PhNuHA-unsplash-1.jpg",
  "https://cnnespanol.cnn.com/wp-content/uploads/2021/06/2F210610151521-splurge-hotel-soneva-jani-super-tease.jpg?quality=100&strip=info",
  "https://static.hosteltur.com/app/public/uploads/img/articles/2024/01/05/L_124129_los-10-mejores-hoteles-del-mundo-segun-tripadvisor.jpg",
  "https://hotelesen.net/wp-content/uploads/2019/12/Hotel-resort-Bellagio-Las-Vegas-elegancia-belleza-fuente-.png",
];
const HomeMainCarousel = () => {
  return (
    <div>
      <HomeCarousel images={url} />
    </div>
  );
};
export default HomeMainCarousel;
