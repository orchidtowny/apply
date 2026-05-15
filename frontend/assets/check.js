window.onload = () => {
    const form = document.getElementById("check-form")
    if (!(form instanceof HTMLFormElement)) return

    form.onsubmit = (ev) => {
        ev.preventDefault()

        const data = new FormData(event.target);
        const username = data.get("username")

        window.fetch(
            "/api/status/"+username,
            { method: "GET" }
        ).then(async (response) => {
            console.log(response)

            if (response.status === 200) {
                const json = await response.json()

                let message
                switch (json?.status) {
                    case '0':
                        message = "Application still awaiting approval."
                        break
                    case '1':
                        message = "Application rejected."
                        break
                    case '2':
                        message = "Application approved."
                        break
                }

                document.getElementById("checked-user").innerHTML = `
                    <p>${message}</p>
                `
            } else if (response.status === 404) {
                document.getElementById("checked-user").innerHTML = `
                    <p>Application not found.</p>
                `
            } else {
                document.getElementById("checked-user").innerHTML = `
                    <p>Something went wrong, try again later.</p>
                `
            }
        })
    }
}