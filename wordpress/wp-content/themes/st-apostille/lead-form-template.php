<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Lead Submission</title>
    <style>
        body { font-family: Arial, sans-serif; text-align: center; margin: 50px; }
        form { max-width: 400px; margin: auto; padding: 20px; border: 1px solid #ccc; border-radius: 5px; background: #f9f9f9; }
        label { display: block; margin: 10px 0 5px; font-weight: bold; }
        input,textarea { width: 100%; padding: 8px; margin-bottom: 10px; border: 1px solid #ccc; border-radius: 3px; }
        button { background: #007bff; color: white; padding: 10px; border: none; cursor: pointer; width: 100%; }
        button:hover { background: #0056b3; }
    </style>
</head>
<body>

    <h2>Lead Submission</h2>
    <form id="leadForm">
        <label for="name">Name *</label>
        <input type="text" id="name" name="name" required>

        <label for="email">Email *</label>
        <input type="email" id="email" name="email" required>

        <label for="phone">Phone Number *</label>
        <input type="tel" id="phone" name="phone" required>

 	<label for="source">Source</label>
        <input type="text" id="source" name="source">

        <label for="message">Message</label>
        <textarea id="message" name="message"></textarea>


        <button type="submit">Submit</button>
    </form>

    <script>
        document.getElementById("leadForm").addEventListener("submit", function(event) {
            event.preventDefault(); // Prevent default form submission

            let formData = {
                name: document.getElementById("name").value,
                email: document.getElementById("email").value,
                phone: document.getElementById("phone").value,
		source: document.getElementById("source").value,
		message: document.getElementById("message").value
            };

            fetch("http://localhost:8181/api/leads", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(formData)
            })
            .then(response => response.json())
            .then(data => {
		if (data.status=="success"){
	 		alert("Form submitted successfully!");
		}else{
			alert(data.message)
		}
                console.log(data);
            })
            .catch(error => {
                alert("Error submitting form!");
                console.error(error);
            });
        });
    </script>

</body>
</html>
