import $axios from '@/utils/axios';

export function getServices() {
    return $axios.get(`/v1/services`);
}

export function getService(name: string) {
    return $axios.get(`/v1/service/${name}`);
}

export function getWebServices() {
    return $axios.get(`/v1/web-services`);
}

export function getAPIGatewayServices() {
    return $axios.get(`/v1/api-gateway-services`);
}

export function getMicroServices() {
    return $axios.get(`/v1/micro-services`);
}
