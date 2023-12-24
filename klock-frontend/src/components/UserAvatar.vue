<template>
  <div
  :class="isUpload ? 'w-[4.5rem] h-[4.5rem] md:w-[6rem] md:h-[6rem]' : ' h-[3rem] w-[3rem] md:h-[3.4rem] md:w-[3.4rem]'"
    class="relative rounded-full"
  >
    <img
      v-if="src || selectedImage"
      class="h-full w-full rounded-full bg-cover bg-center bg-no-repeat"
      :src="src || previewImage"
      :alt="alt"
    />
    <div
      v-else
      class="absolute left-0 top-0 grid h-full w-full place-items-center rounded-full border-2 border-slate-400 bg-slate-300"
    >
      <UserIcon />
    </div>
    <label v-if="isUpload" for="upload-avatar">
      <div
      :class="isUpload ? 'h-6.5 w-6.5 md:h-8 md:w-8' : 'h-5 w-5'"
        class="absolute bottom-2 -right-2 z-10 flex cursor-pointer items-center justify-center rounded-full border border-[#667085] bg-white"
      >
        <img
          src="../assets/camera.svg"
          alt="camera"
          :class="isUpload ? 'scale-90' : 'scale-75'"
          class="transform"
        />
        <input
          type="file"
          hidden
          id="upload-avatar"
          accept=".svg, .png, .jpg, .jpeg, .gif"
          @change="handlepAvatarUpload"
        />
      </div>
    </label>
  </div>
</template>

<script setup lang="ts">
import { errorApiRequest, isValidImage, successApiRequest } from "@/lib/helperFunctions";
import UserIcon from "../assets/UserIcon.vue";
import { ref } from "vue";
import type { ErrorObject } from "@/types";
import { updateAvatar } from "@/api/user";

const emit = defineEmits(["uploadImage"]);

defineProps({
  src: {
    type: String,
  },
  alt: {
    type: String,
    required: true,
  },
  isUpload: {
    type: Boolean,
    default: false,
  },
});

const previewImage = ref("");
const selectedImage = ref<File | null>(null);

const handlepAvatarUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const files = target.files as FileList;
  processFiles(files);
};

const processFiles = (files: FileList) => {
  const file = files[0];

  isValidImage(file)
    .then((validatedFile) => {
      selectedImage.value = validatedFile;
      previewImage.value = URL.createObjectURL(validatedFile);
      updateAvatar(validatedFile)
      .then((res) => {
        console.log(res);
        // successApiRequest(res)
      })
      .catch(err => {
        console.error(err);
        errorApiRequest(err);
      })
    })
    .catch((error: ErrorObject) => {
      errorApiRequest(error);
    });
};
</script>
