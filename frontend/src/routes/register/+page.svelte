<script lang="ts">
  let email = '';
  let username = '';
  let password = '';
  let passwordRepeat = '';

  async function register() {
    if (password !== passwordRepeat) {
      alert("Passwords do not match!");
      return;
    }

    // This is where you handle the registration logic.
    // The code below sends the registration request to your backend.
    try {
      const response = await fetch('http://localhost:8080/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, username, password }) // Changed to send email and password
      });

      // This is where you handle the response from the server.
      if (response.ok) {
        // Success! For example, you can redirect the user to the login page.
        alert('Registration successful!');
        window.location.href = '/login';
      } else {
        // Handle errors, e.g., display a message to the user.
        const errorData = await response.json();
        alert(`Registration failed: ${errorData.message}`);
      }
    } catch (error) {
      console.error('An error occurred during registration:', error);
      alert('An error occurred. Please try again.');
    }
  }
</script>

<form on:submit|preventDefault={register}>
  <div class="container">
    <h1>Register</h1>
    <p>Please fill in this form to create an account.</p>
    <hr>

    <label for="email"><b>Email</b></label>
    <input type="text" placeholder="Enter Email" name="email" id="email" bind:value={email} required>

    <label for="username"><b>Username</b></label>
    <input type="text" placeholder="Enter Username" name="username" id="username" bind:value={username} required>        

    <label for="psw"><b>Password</b></label>
    <input type="password" placeholder="Enter Password" name="psw" id="psw" bind:value={password} required>

    <label for="psw-repeat"><b>Repeat Password</b></label>
    <input type="password" placeholder="Repeat Password" name="psw-repeat" id="psw-repeat" bind:value={passwordRepeat} required>
    <hr>

    <p>By creating an account you agree to our <a href="#">Terms & Privacy</a>.</p>
    <button type="submit" class="registerbtn">Register</button>
  </div>

  <div class="container signin">
    <p>Already have an account? <a href="login">Sign in</a>.</p>
  </div>
</form> 