import { useState } from "react";

function Sign_up(){
    const [user, set_book] = useState({
            user_name: "",
            real_name: "",
            password: "",
        });

        const handleChange = (e) => {
            const { name, value } = e.target;
            set_book({
                ...user,
                [name]: value,
            });
        };

        const send_form = async (e) => {
            e.preventDefault();
            const url = "http://localhost:8080/user";
            try {
                const response = await fetch(url, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        user_id: user.user_name,
                        user_name: user.real_name,
                        user_password: user.password,
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
                    alert("the user was saved");
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
                    <h2 className="text-3xl font-bold text-gray-800 mb-8 text-center">Sign Up</h2>
                    <form onSubmit={send_form}>
                        <div className="mb-6">
                            <label className="block text-sm font-semibold text-gray-700 mb-2">User name</label>
                            {form_input("text","title",user.user_name,handleChange,"EJ: El Jose 123")}
                        </div>

                        <div className="mb-6">
                            <label className="block text-sm font-semibold text-gray-700 mb-2">Real Name</label>
                            {form_input("text","id",user.real_name,handleChange,"EJ: Jose Perez")}
                        </div>

                        <div className="mb-6">
                            <label className="block text-sm font-semibold text-gray-700 mb-2">Password</label>
                            {form_input("password","author_name",user.password,handleChange,"EJ: MyPassword123")}
                        </div>

                        <button 
                            type="submit"
                            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg"
                        >
                            sign up
                        </button>
                    </form>
                </div>
            </section>
        );
    }

export default Sign_up