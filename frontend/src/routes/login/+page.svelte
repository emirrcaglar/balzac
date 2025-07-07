<script>
  let username = '';
  let password = '';

  async function login() {
    try {
      const response = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    if (response.ok) {
      alert('Login successful!');
      window.location.href = '/';
    } else {
        const errorData = await response.json();
        alert(`Login failed: ${errorData.message}`);      
    }

    const data = await response.json();
    // handle response...
    } catch (error) {
      console.error('An error occurred during login:', error);
      alert('An error occurred. Please try again.');      
    }
 
  }
</script>

<form on:submit|preventDefault={login}>
  <div class="container">
    <h1>Login</h1>
    <p>Please fill in this form to log in.</p>
    <hr>

    <label for="username"><b>Username</b></label>
    <input type="text" placeholder="Enter Username" name="username" id="username" required>    

    <label for="psw"><b>Password</b></label>
    <input type="password" placeholder="Enter Password" name="psw" id="psw" required>

    <button type="submit" class="loginbtn">Login</button>
  </div>

  <div class="container login">
    <p>Don't have an account? <a href="register">Register</a>.</p>
  </div>
</form> 