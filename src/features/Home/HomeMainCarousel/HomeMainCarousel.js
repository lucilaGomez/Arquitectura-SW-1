import HomeCarousel from "../../../components/HomeCarousel/HomeCarousel";

const images = [
  "https://hips.hearstapps.com/hmg-prod/images/oliver-vegas-1600328764.jpg?crop=1.00xw:1.00xh;0,0&resize=768:*",
  "https://img.freepik.com/foto-gratis/lobo-colorido-fondo-negro_1340-40203.jpg",
  "https://dus6dayednven.cloudfront.net/app/uploads/2020/12/JEAN-PAUL-DE-LA-HARPE.jpg",
  "https://denomades.s3.us-west-2.amazonaws.com/blog/wp-content/uploads/2016/05/12115951/como-sacar-buenas-fotos-denomades-5.jpg",
];
const HomeMainCarousel = () => {
  return (
    <div>
      <HomeCarousel images={images} />
    </div>
  );
};
export default HomeMainCarousel;
