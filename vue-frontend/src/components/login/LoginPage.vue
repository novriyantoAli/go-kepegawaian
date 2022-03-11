<template>
  <section>
    <v-container class="fill-height" fluid>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="4">
          <v-card class="mx-auto">
            <v-toolbar color="primary" dark flat>
              <v-toolbar-title>Login</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form ref="form" v-model="valid">
                <v-text-field
                  label="username"
                  v-model="username"
                  prepend-icon="email"
                  type="text"
                  :rules="numberRules"
                  required>
                </v-text-field>

                <v-text-field
                  label="Password"
                  v-model="password"
                  prepend-icon="lock"
                  type="password"
                  :rules="passwordRules"
                  required
                >
                </v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <div class="flex-grow-1"></div>
              <v-btn color="primary" :disabled="!valid" @click="doLogin">
                <span class="spinner-border spinner-border-sm" v-show="loading"></span> Login
              </v-btn>
            </v-card-actions>
          </v-card>
          <v-snackbar v-model="snackbar" bottom>
            {{ error_message }}
          </v-snackbar>
        </v-col>
      </v-row>
    </v-container>
  </section>
</template>
 
<script>
import { authenticationService } from '@/_services';
export default {
  props: {
    source: String,
  },
  data: () => ({
    drawer: null,
    snackbar: false,
    valid: true,
    loading: false,
    username: "",
    numberRules: [
      (v) => v.length > 0 || "This field may not be empty",
      (v) => v > 0 || "The value must be greater than zero",
    ],
    password: "",
    passwordRules: [
      (v) => !!v || "Password is required",
      (v) =>
        (v && v.length >= 6) || "Password must be larger than 6 characters",
    ],
    result: null,
    error_message: "",
  }),
  created () {
    // redirect to home if already logged in
    if (authenticationService.currentUserValue) { 
      return this.$router.push({ name: "dashboard" });
    }
    // get return url from route parameters or default to '/'
    this.returnUrl = this.$route.query.returnUrl || '/';
  },
  methods: {
    reset() {
      this.$refs.form.reset();
    },
    doLogin() {
      this.loading = true
      authenticationService.login(this.username, this.password).then(
        user => {
          this.error_message = "wellcome "+user.username;
          this.$router.push(this.returnUrl);
          this.loading = false
        }, 
        error => {
          console.log(error);
          this.loading = false;
          this.error_message = error;
          this.snackbar = true;
        }
      );
    },
  },
};
</script>