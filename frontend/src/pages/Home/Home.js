import HeroBanner from "../../features/Home/HeroBanner/HeroBanner"
import HomeMainCarousel from "../../features/Home/HomeMainCarousel/HomeMainCarousel"
import HomeOffers from "../../features/Home/HomeOffers/HomeOffers"
import MainLayout from "../../layouts/MainLayout/MainLayout"

const Home = () => {
  return (
    <div>
      <MainLayout overflow="hidden">
        <HeroBanner></HeroBanner>
        {/* <HomeOffers></HomeOffers> */}
        <HomeMainCarousel></HomeMainCarousel>
      </MainLayout>
    </div>
  )
}
export default Home