import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import "./MainLayout.css";
const MainLayout = ({children, overflow = 'visible'}) => {
  return (
    <div>
        <Header></Header>
        <main>
                {children}
            </main>
            <Footer/>
    </div>
  )
}
export default MainLayout