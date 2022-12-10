import EmbedPageExample from "../pages/EmbedPageExample";
import EmbedPageWithID from "../pages/EmbedPageExample/EmbedPageWithID";
import Home from "../pages/Home";
import { Routes, Route } from "react-router-dom";
import NavBar from "../components/NavBar";
import Login from "../pages/Login"
import PageNotFound from "../pages/PageNotFound";
import ContactUs from "../pages/ContactUs";
import Signup from "../pages/Signup";
import MockExample from "../pages/MockExample";
import { useAuthState } from "../hook/auth";
export default function Router() {
    const auth = useAuthState()
    return (
        <Routes>
            <Route element={<NavBar user={auth.user} />}>
                <Route path="/" element={<Home />} />
                <Route path="/embed_page_example" element={<EmbedPageExample />}>
                    <Route path=":id" element={<EmbedPageWithID />} />
                </Route>
                <Route path="/contact_us" element={<ContactUs />} />
                <Route path="/mock_example" element={<MockExample />} />
                <Route path="/login" element={<Login />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/*" element={<PageNotFound />} />
            </Route>
        </Routes>
    )
}