import ErrorPage from "../../features/Error404Page/ErrorPage";
import MainLayout from "../../layouts/MainLayout/MainLayout";
import './Error404Page.css';

const Error404Page = () => {
    return (
        <MainLayout overflow="hidden">
            <section className="error404Page-section">
                <ErrorPage />
            </section>
        </MainLayout>
    )
}
export default Error404Page