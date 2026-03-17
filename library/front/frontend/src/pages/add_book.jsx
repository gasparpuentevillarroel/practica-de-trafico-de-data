import { useState } from "react";
import form_input from "../components/input";

function Add_book() {
    const [book, set_book] = useState({
        id: "",
        title: "",
        author_name: "",
        author_id: "",
        year_publication: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        set_book({
            ...book,
            [name]: value,
        });
    };

    const send_form = async (e) => {
        e.preventDefault();
        const url = "http://localhost:8080/books";
        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    id: book.id,
                    title: book.title,
                    author_name: book.author_name,
                    author_id: Number(book.author_id),
                    year_publication: Number(book.year_publication),
                }),
            });
            if (response.ok) {
                await response.json();
                set_book({
                    id: "",
                    title: "",
                    author_name: "",
                    author_id: "",
                    year_publication: "",
                });
                alert("the book was saved");
            } else {
                const data = await response.json().catch(() => ({}));
                throw new Error(data.error || `Error del servidor (${response.status})`);
            }
        } catch (error) {
            console.error("conection error:", error);
            alert(error.message || "Error de conexión");
        }
    };

    return (
        <section className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-50 to-gray-100 p-4">
            <div className="max-w-md mx-auto p-8 bg-white rounded-2xl shadow-xl">
                <h2 className="text-3xl font-bold text-gray-800 mb-8 text-center">Cargar libro</h2>
                <form onSubmit={send_form}>
                    <div className="mb-6">
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Título</label>
                        {form_input("text","title",book.title,handleChange,"EJ: El Tunel")}
                    </div>

                    <div className="mb-6">
                        <label className="block text-sm font-semibold text-gray-700 mb-2">ID</label>
                        {form_input("text","id",book.id,handleChange,"EJ: 8034")}
                    </div>

                    <div className="mb-6">
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Autor</label>
                        {form_input("text","author_name",book.author_name,handleChange,"EJ: Ernesto Sabato")}
                    </div>

                    <div className="mb-6">
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Autor ID</label>
                        {form_input("number","author_id",book.author_id,handleChange,"Ej: 1")}
                    </div>

                    <div className="mb-8">
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Año</label>
                        {form_input("number","year_publication",book.year_publication,handleChange,"Ej: 1948")}
                    </div>

                    <button 
                        type="submit"
                        className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg"
                    >
                        Guardar en Biblioteca
                    </button>
                </form>
            </div>
        </section>
    );
}

export default Add_book;