window.onload = () => {
    const form = document.getElementById("form")
    if (!(form instanceof HTMLFormElement)) return

    const alert = document.getElementById("alert")
    if (!(form instanceof HTMLDivElement)) return

    form.onsubmit = (ev) => {
        ev.preventDefault()

        window.fetch(
            "/api/apply",
            {
                method: "POST",
                body: JSON.stringify({
                    username: "",
                    age: 0,
                    where_did_you_find_the_server: "",
                    bio: ""
                })
            }
        ).then((response) => {
            console.log(response)
        })
    }
}