import HotelSearcher from "../../features/HotelFinder/HotelSearcher/HotelSearcher"
import MainLayout from "../../layouts/MainLayout/MainLayout"

const HotelFinder = () => {
  return (
    <div>
      <MainLayout overflow="hidden">
        <HotelSearcher></HotelSearcher>
      </MainLayout>
    </div>
  )
}
export default HotelFinder