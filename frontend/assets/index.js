window.onload = () => {
    const form = document.getElementById("form")
    if (!(form instanceof HTMLFormElement)) {
        console.log("error with form, fail at !instanceof")
        return
    }

    form.onsubmit = (ev) => {
        ev.preventDefault()

        const data = new FormData(event.target);
        const username = data.get("username")
        const age = data.get("age")
        const whereDidYouFindTheServer = data.get("where_did_you_find_the_server")
        const bio = data.get("bio")

        window.fetch(
            "/api/apply",
            {
                method: "POST",
                body: JSON.stringify({
                    username: username,
                    age: Number(age),
                    where_did_you_find_the_server: whereDidYouFindTheServer,
                    bio: bio
                })
            }
        ).then(async (response) => {
            console.log(response)

            let msg = response.statusText

            try {
                msg = await response.text()
            } catch (_) {}

            if (response.ok) window.location.replace(`/status?username=${username}`)
            else alert(msg + "\n\n Try again later or contact an admin.")
        })
    }
}