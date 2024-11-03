const minc = document.querySelector(".minc")
const maxc = document.querySelector(".maxc")
const minvc = document.querySelector(".fromCreation")
const maxvc = document.querySelector(".toCreation")
const trackc = document.querySelector(".trackCreation")

function updateTrack() {
    const minValue = parseInt(minc.value);
    const maxValue = parseInt(maxc.value);
    const minPercent = ((minValue - 1958) / (2015 - 1958)) * 100;
    const maxPercent = ((maxValue - 1958) / (2015 - 1958)) * 100;
    
    trackc.style.background = `linear-gradient(to right, lightblue ${minPercent}%, #cdcdcd ${minPercent}%, #cdcdcd ${maxPercent}%, lightblue ${maxPercent}%)`;
}

minc.addEventListener('input', (e) => {
    if (minc.value > maxc.value) {
        minc.value = maxcc.value
    }
    minvc.textContent = minc.value
    updateTrack();
})

maxc.addEventListener('input', () => {
    if (minc.value > maxc.value) {
        maxc.value = minc.value
    }
    maxvc.textContent = maxc.value
    updateTrack();
})

minvc.textContent = minc.value;
maxvc.textContent = maxc.value;
updateTrack();


const minf = document.querySelector(".minf")
const maxf = document.querySelector(".maxf")
const minvf = document.querySelector(".fromAlbum")
const maxvf = document.querySelector(".toAlbum")
const trackf = document.querySelector(".trackAlbum")
function updateTrackf() {
    const minValue = parseInt(minf.value);
    const maxValue = parseInt(maxf.value);
    const minPercent = ((minValue - 1963) / (2018 - 1963)) * 100;
    const maxPercent = ((maxValue - 1963) / (2018 - 1963)) * 100;
    
    trackf.style.background = `linear-gradient(to right, lightblue ${minPercent}%, #cdcdcd ${minPercent}%, #cdcdcd ${maxPercent}%, lightblue ${maxPercent}%)`;
}

minf.addEventListener('input', (e) => {
    if (minf.value > maxf.value) {
        minf.value = maxf.value
    }
    minvf.textContent = minf.value
    updateTrackf();
})

maxf.addEventListener('input', () => {
    if (minf.value > maxf.value) {
        maxf.value = minf.value
    }
    maxvf.textContent = maxf.value
    updateTrackf();
})
minvf.textContent = minf.value;
maxvf.textContent = maxf.value;
updateTrackf();