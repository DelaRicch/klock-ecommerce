<template>
  <div class="flex flex-col gap-1 w-full">
    <div class="flex gap-1 items-center">
    <label v-if="type !== 'file' " :for="id" class="text-[#1D2939] font-medium line-clamp-1">
      {{ label }}
    </label>
      <span v-if="requiredTag" class="text-red-600 font-semibold">*</span>
    </div>
    <div v-if="type !== 'file'" :class="{'h-[2.9rem]': textType === 'input',

    'h-[7rem]': textType === 'textarea',
    className
    }" class="relative">
      <input
          v-if="textType == 'input'"
          @input.prevent="$emit('updateValue', ($event.target as HTMLInputElement).value.trim());"
             :type="inputType"
          :id="id"
          :hidden="type === 'file'"
          :placeholder="placeholder"
             class="placeholder:text-[#667085] rounded-lg h-full border border-[#667085] w-full px-3 outline-none outline-offset-4 hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] hover:border-[#0408E7] pr-8"
      />

      <textarea
          v-if="textType === 'textarea'"
          @input.prevent="$emit('updateValue', ($event.target as HTMLInputElement).value.trim())"
          :id="id"
          v-model="inputString"
          :placeholder="placeholder"
          :maxlength="maxCharacters"
          class="resize-none rounded-lg h-full border border-[#667085] w-full px-3 outline-none outline-offset-4 hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] hover:border-[#0408E7] pr-8"
      ></textarea>
      <div v-if="textType === 'textarea'"
           class="text-[0.75rem] flex items-center justify-end text-slate-400 pr-1">{{characterCount}}/{{
          maxCharacters }}</div>

      <button
          type="button"
          v-if="type === 'password'"
               @click="handleToggleIsHidden"
              class="absolute right-2 top-[50%] translate-y-[-50%] rounded-lg outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]">
        <EyeViewComponent :is-hidden="isHidden" />
      </button>
    </div>
      <label v-if="type === 'file'" :for="id">
        <div class="flex gap-1 flex-col w-full">
         <span class="text-[#1D2939] font-medium line-clamp-1"> {{label}}</span>
          <input
              v-if="textType == 'input'"
              :type="inputType"
              :id="id"
              hidden
              @change="handleFileUpload"
              accept=".svg, .png, .jpg, .jpeg, .gif"
          />
<div
    @dragover.prevent
    @dragenter.prevent
    @drop="handleDrop"
    :class="{'h-[15rem]' : !selectedImage}"
    class="cursor-pointer placeholder:text-[#667085] rounded-lg max-h-[18rem] border border-[#667085] w-full outline-none outline-offset-4 hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC] hover:border-[#0408E7] flex items-center justify-center overflow-clip"
>
  <img
      v-if="selectedImage"
      :src="previewImage"
      alt="selected cover image"
      class="w-full max-h-full object-cover"
  >
<div
    v-if="!selectedImage"
    class="w-9/12 h-4/6 flex flex-col gap-2 items-center justify-center">
  <img class="w-[46px]" :src="FileUpload" alt="file-upload">
  <div class="flex flex-col text-sm text-[#667085]">
  <span class="line-clamp-1"><span class="text-blue-600">Click to upload </span>or drag and drop</span>
  <span class="text-[0.7rem] line-clamp-1">SVG, PNG, JPG or GIF (max. 800x400px)</span>
  </div>
</div>
</div>
        </div>
      </label>
    <span v-if="!!error"
          class="text-sm text-red-600">{{ error }}</span>
  </div>


</template>


<script setup lang="ts">
import EyeViewComponent from "./EyeViewComponent.vue";
import FileUpload from '@/assets/file-upload.svg'
import {computed, ref, watch} from "vue";
import {errorApiRequest, isValidImage, processFiles} from "@/lib/helperFunctions.ts";
import {ErrorObject} from "@/types";

const { label, placeholder, type, id, requiredTag } = defineProps({
  label: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    required: false,
  },
  type: {
    type: String,
    default: 'text',
  },
  textType: {
    type: String,
    default: 'input',
    validator: (value: string) => ['input', 'textarea'].includes(value),
  },
  id: {
    type: String,
    required: true,
  },
  requiredTag: {
    type: Boolean,
    default: false,
  },
  error: {
    type: String,
    default: '',
  },
  maxCharacters: {
    type: Number,
    default: 100,
  },
  className: {
    type: String,
    required: false,
  }
})

defineEmits(['updateValue'])

const isHidden = ref(false);
const inputType = ref(type);
const inputString = ref<string>('');
const uploading = ref(false);
const selectedImage = ref<File | null>(null);
const previewImage = ref('');


const handleToggleIsHidden = () => isHidden.value = !isHidden.value;

const characterCount = computed(() => inputString.value.trim().length);

const handleDrop = (event: DragEvent) => {
  event.preventDefault();
  const droppedFiles = (event.dataTransfer as DataTransfer).files;
  processFiles(droppedFiles);
}

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const files = target.files as FileList;
  console.log(files, "files");
  processFiles(files);
}

const processFiles = (files: FileList) => {
  const file = files[0];

  isValidImage(file)
      .then((validatedFile) => {
        console.log(validatedFile, "Selected file");
        // Further processing with the validated file
        // Assuming selectedImage and previewImage are ref objects
        selectedImage.value = validatedFile;
        previewImage.value = URL.createObjectURL(validatedFile);
      })
      .catch((error: ErrorObject) => {
        errorApiRequest(error)
      });
};

watch(isHidden, (value) => {
  inputType.value = value ? "text" : type;
});
</script>