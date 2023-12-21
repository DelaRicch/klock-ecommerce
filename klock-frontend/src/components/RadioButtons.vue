<template>
    <div :class="{ 'flex-col gap-4': orientation === 'vertical', 'flex-row gap-8': orientation === 'horizontal' }" class="flex">
        <label v-for="radio in options" :key="radio" :for="radio" class="flex items-center gap-1.5 cursor-pointer text-sm">
            <input type="radio" name="radio-group" :id="radio" class="cursor-pointer text-black p-0 bg-white border-gray-300 outline-none ring-0 ring-transparent rounded-full form-input" :value="radio" @change="handleSelectRadio(radio)" />
            <span class="capitalize">{{radio}}</span>
        </label>
    </div>
</template>
<script lang="ts" setup>
import { type PropType } from 'vue';


const emit = defineEmits(['update:modelValue']);

defineProps({
    options: {
        type: Array as PropType<string[]>,
        required: true
    },
    orientation: {
        type: String,
        required: false,
        validator: (value: string) => {
            return ['horizontal', 'vertical'].includes(value);
        },
        default: 'horizontal'
    }
})

const handleSelectRadio = (value: string) => {
    emit('update:modelValue', value);
}
</script>