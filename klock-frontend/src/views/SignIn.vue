<template>
  <section :class="isDesktop && 'flex h-screen items-center justify-center'">
    <div
      v-if="!isDesktop"
      class="relative -z-10 h-[10.5rem] w-full bg-auth-bg bg-cover bg-center bg-no-repeat"
    >
      <!--      Logo and title for Mobile screen-->
      <div
        class="z-20 flex h-full w-full flex-col items-center justify-center gap-8"
      >
        <LabelLogo
          :text-color="textColor"
          :inner-color="innerColor"
          :outer-color="outerColor"
        />
        <SignUpLoginTitle :title="'Sign In'" :is-desktop="isDesktop" />
      </div>

      <!--Black overlay-->
      <div class="absolute inset-0 -z-10 bg-black opacity-50"></div>
    </div>
    <div
      class="mx-auto mt-10 flex w-[95%] max-w-[38rem] flex-col gap-5 rounded-lg bg-[#FBFBFB] px-2.5 py-6 sm:w-[90%] sm:px-[2.5rem] md:px-[4.13rem]"
    >
      <!-- Logo and title for Desktop screen-->
      <div
        v-if="isDesktop"
        class="z-20 flex h-full w-full flex-col items-center justify-center gap-8"
      >
        <LabelLogo
          :text-color="textColor"
          :inner-color="innerColor"
          :outer-color="outerColor"
        />
        <SignUpLoginTitle :title="'Sign In'" :is-desktop="isDesktop" />
      </div>

      <form @submit.prevent="handleSubmit" class="flex flex-col gap-5">
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
          @update-value="handleValidateInput($event, 'password')"
          :error="formData.password.error"
          :id="'password'"
          :placeholder="'Enter your password'"
          :label="'Password'"
          type="password"
        />

        <div class="flex items-center justify-between">
          <CheckboxComp
            :id="'sign-up'"
            :label="'Remember Me'"
            :error="formData.remember.error"
            @update-value="handleValidateInput($event, 'remember')"
          />
          <router-link
            to="/terms-and-conditions"
            class="p-[2px] text-[0.8rem] text-[#181CF9] outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
          >
            Forgot Password?
          </router-link>
        </div>

        <ButtonComponent
          :color="'white'"
          :background-color="'#1D2939'"
          :type="'submit'"
          :label="'sign In'"
          :is-disabled="!isValidForm"
          uppercase
        />
      </form>
      <div class="mt-2 flex flex-col gap-8">
        <span class="text-center text-[#1D2939]">OR</span>

        <!-- <div class="flex items-center justify-center gap-[1.8rem]">
          <SocialAuthButton authType="Facebook" />
          <SocialAuthButton authType="Google" />
        </div> -->
        <div class="flex items-center justify-center gap-1">
          <p class="text-[0.8rem] text-black/70">Don't have an account?</p>
          <router-link
            to="/sign-up"
            class="p-[2px] text-[0.8rem] text-[#181CF9] outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
            >Sign Up</router-link
          >
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, reactive } from "vue";

import { useRouter } from "vue-router";
import { emailRegex } from "@/schema/ValidationSchema";
import ButtonComponent from "@/components/ButtonComponent.vue";
import CheckboxComp from "@/components/CheckboxComp.vue";
import InputField from "@/components/InputField.vue";
import LabelLogo from "@/assets/LabelLogo.vue";
import SignUpLoginTitle from "@/components/SignUpLoginTitle.vue";
import { innerColor, isDesktop, outerColor, textColor } from "@/stores/resuableState";
import { getUserProfile, loginUser } from "@/api/user";
import { errorApiRequest, handleUserProfile, successApiRequest } from "@/lib/helperFunctions";
import type { formFields } from "@/types";
const router = useRouter();

// Function to update colors based on screen size
const updateColors = () => {
  if (window.innerWidth < 768) {
    isDesktop.value = false;
    // Set mobile colors
    textColor.value = "white";
    innerColor.value = "white";
    outerColor.value = "#1D2939";
  } else {
    isDesktop.value = true;
    // Set desktop colors
    textColor.value = "#1D2939";
    innerColor.value = "#1D2939";
    outerColor.value = "white";
  }
};

const formData: formFields = reactive({
  email: {
    value: "",
    error: "",
  },
  password: {
    value: "",
    error: "",
  },
  remember: {
    value: "false",
    error: "",
  },
});

const handleSubmit = async () => {
  const formDataValues: Record<string, string> = {};
  for (const field in formData) {
    formDataValues[field] = (formData[field] as { value: string }).value;
  }

  loginUser(formDataValues)
    .then((res) => {
      successApiRequest(res);
    })
    .then(() => {
      getUserProfile().then((res) => {
        handleUserProfile(res.user);
        router.push("/dashboard");
      });
    })
    .catch((err) => {
      errorApiRequest(err);
    });
};

onMounted(() => {
  updateColors();
  window.addEventListener("resize", updateColors);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", updateColors);
});

const isValidForm = computed(() => {
  return Object.values(formData).every(
    (field) => field.value !== "" && !field.error,
  );
});

const handleValidateInput = (value: string, field: string) => {
  if (field === "email" && !emailRegex.test(value)) {
    formData[field].error = "Invalid email format";
  } else {
    formData[field].error = "";
  }

  formData[field].value = value.toString();
};
</script>
