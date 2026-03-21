
import { BrowserRouter, Route,Routes } from "react-router-dom";
import Sign_up from './pages/sign_up.jsx';
import Add_book from "./pages/add_book.jsx";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/add_user" element={<Sign_up />}/>
                <Route path="/add_book" element={<Add_book/>}/>
            </Routes>
        </BrowserRouter>
    )
}
export default App;
