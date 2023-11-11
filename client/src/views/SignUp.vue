<template>
  <section :class="isDesktop && 'flex items-center justify-center h-screen'">
    <div v-if="!isDesktop"
         class="bg-auth-bg relative w-full h-[10.5rem] bg-center bg-no-repeat bg-cover -z-10">

      <!--      Logo and title for Mobile screen-->
      <div class="z-20 w-full h-full flex flex-col gap-8 justify-center items-center">
        <LabelLogo :text-color="textColor" :inner-color="innerColor" :outer-color="outerColor" />
        <SignUpLoginTitle :title="'Sign Up'" :is-desktop="isDesktop" />
      </div>

      <!--Black overlay-->
      <div class="absolute inset-0 bg-black opacity-50 -z-10"></div>

    </div>
    <div
        class="mt-10 w-[95%] sm:w-[90%] max-w-[38rem] mx-auto flex flex-col gap-5 bg-[#FBFBFB] rounded-lg px-2.5 sm:px-[2.5rem] md:px-[4.13rem] py-6">

      <!-- Logo and title for Desktop screen-->
      <div v-if="isDesktop" class="z-20 w-full h-full flex flex-col gap-8 justify-center items-center">
        <LabelLogo :text-color="textColor" :inner-color="innerColor" :outer-color="outerColor" />
        <SignUpLoginTitle :title="'Sign Up'" :is-desktop="isDesktop" />
      </div>

      <form @submit.prevent="handleSubmit" class="flex flex-col gap-5">
        <InputField
            @update-value="handleValidateInput($event, 'name')"
            :error="formData.name.error"
            required-tag
            :id="'name'"
            :placeholder="'Enter your full name'"
            :label="'Name'"
        />

        <InputField
            required-tag
            @update-value="handleValidateInput($event, 'email')"
            :error="formData.email.error"
            :id="'email'"
            :placeholder="'Enter your email'"
            :label="'Email'"
            type="email"
        />
        <InputField
            required-tag
            @update-value="handleValidateInput($event,'password')"
            :error="formData.password.error"
            :id="'password'"
            :placeholder="'Enter your password'"
            :label="'Password'"
            type="password"
            :isPassword="true"

        />
        <InputField
            required-tag
            @update-value="handleValidateInput($event, 'confirm_password')"
            :error="formData.confirm_password.error"
            :id="'confirm_password'"
            :placeholder="'Confirm your password'"
            :label="'Confirm Password'"
            type="password"
            :isPassword="true" />

        <div class="flex items-center gap-1">
          <CheckboxComp
              :id="'sign-up'"
              :label="'I agree the'"
              :error="formData.agreeToTerms.error"
              @update-value="handleValidateInput($event, 'agreeToTerms')"
          />
          <router-link to="/terms-and-conditions"
                       class="text-[#181CF9] p-[2px] text-[0.8rem] outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]">Terms and
            Conditions</router-link>
        </div>

        <ButtonComponent :color="'white'"
                         :background-color="'#1D2939'"
                         :type="'submit'"
                         :label="'sign up'"
                         :is-disabled="!isValidForm"
                         uppercase />
      </form>
      <div class="flex flex-col gap-8 mt-2">
        <span class="text-[#1D2939] text-center">OR</span>

        <div class="flex items-center gap-[1.8rem] justify-center">
          <SocialAuthButton authType="Facebook" />
          <SocialAuthButton authType="Google" />
        </div>
        <div class="flex items-center gap-1 justify-center">
          <p class="text-[0.8rem] text-[#1D2939]">Already have an account?</p>
          <router-link to="/sign-in"
                       class="text-[#181CF9] p-[2px] text-[0.8rem] outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]">Log In</router-link>
        </div>
      </div>
    </div>
  </section>
  <Snackbar
      :show-snackbar="showSnackbar"
      :success="success"
      :title="success ? 'success' : 'error' "
            :message="apiResponse" />
</template>


<script setup lang="ts">
import InputField from "../components/InputField.vue";
import LabelLogo from "../assets/LabelLogo.vue";
import {computed, onBeforeUnmount, onMounted, reactive} from "vue";
import SignUpLoginTitle from "../components/SignUpLoginTitle.vue";
import CheckboxComp from "../components/CheckboxComp.vue";
import ButtonComponent from "../components/ButtonComponent.vue";
import SocialAuthButton from "../components/SocialAuthButton.vue";
import { signUpFormFields} from "@/types";
import {emailRegex, passwordRegex} from "../schema/ValidationSchema.ts";
import {registerUser} from "../api/user.ts";

import {
  errorApiRequest,
  successApiRequest
} from "../lib/helperFunctions.ts";
import Snackbar from "../components/Snackbar.vue";
import {useUserStore} from "../store/store.ts";
import {
  apiResponse,
  innerColor,
  isDesktop,
  outerColor,
  showSnackbar, success,
  textColor
} from "../store/resuableState.ts";

const userStore = useUserStore();

// Function to update colors based on screen size
const updateColors = () => {
  if (window.innerWidth < 768) {
isDesktop.value = false;
    // Set mobile colors
    textColor.value = 'white';
    innerColor.value = 'white';
    outerColor.value = '#1D2939';
  } else {
isDesktop.value = true;
    // Set desktop colors
    textColor.value = '#1D2939';
    innerColor.value = 'black';
    outerColor.value = 'white';
  }
};

const formData: signUpFormFields = reactive({
  name: {
    value: '',
    error: '',
  },
  email: {
    value: '',
    error: '',
  },
  password: {
    value: '',
    error: '',
  },
  confirm_password: {
    value: '',
    error: '',
  },
  agreeToTerms: {
    value: 'false',
    error: '',
  },
});

const handleSubmit = async () => {
  const formDataValues: Record<string, string> = {};
  for (const field in formData) {
    formDataValues[field] = (formData[field] as { value: string }).value;
  }

registerUser(formDataValues)
    .then((res) => {
     successApiRequest(res);

    })
    .catch((err) => {
      errorApiRequest(err);
    });

};

onMounted(() => {
  updateColors();
  window.addEventListener('resize', updateColors);
});

onBeforeUnmount(() => {
  window.removeEventListener('resize', updateColors);
});

const isValidForm = computed(() => {
  return Object.values(formData).every((field) => (field.value !== '' && field.value !== 'false')
      &&
      !field.error)
});


const handleValidateInput = (value: string, field: string) => {
  if (field === 'agreeToTerms' && value.toString() !== 'true') {
    formData[field].error = 'This field is required';
  } else if (field !== 'agreeToTerms' && !value.trim()) {
    formData[field].error = 'This field is required';
  } else if (field === 'email' && !emailRegex.test(value)) {
    formData[field].error = 'Invalid email format';
  } else if (field === 'password' && !passwordRegex.test(value)) {
    formData[field].error = 'Password must be at least 8 characters and include one uppercase letter, one lowercase letter, and one digit';
  } else if (field === 'confirm_password') {
    const password = formData['password'].value;
    if (value !== password) {
      formData[field].error = 'Passwords do not match';
    } else {
      formData[field].error = '';
    }
  } else {
    formData[field].error = '';
  }

  formData[field].value = value.toString();
};

</script>
