export interface logEntry {
    id: number
    // trackedDeviceId: number
    trackedDevice?: trackedDevice
    log_time: Date
    sensorValue: number
}

export interface deviceType {
    id: number
    name: string
}


export interface trackedDevice {
    id: number
    // deviceTypeId: number
    deviceType?: deviceType
    location: string
}
    

export interface sensor {
    id: string
    // trackedDeviceId: number
    trackedDevice?: trackedDevice
    batteryStatus: number
}
    