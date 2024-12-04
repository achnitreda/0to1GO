let debounceTimer;

async function fetchSuggestions(query) {
    // if after trimming query is empty string
    if (!query.trim()) {
        suggestionsContainer.style.display = 'none';
        return;
    }

    try {
        const resp = await fetch(`/search/?q=${encodeURIComponent(query.trim())}`)
        const suggestions = await resp.json()
        if (suggestions && suggestions.length > 0) {
            displaySuggestions(suggestions);
        } else {
            suggestionsContainer.style.display = 'none';
        }

    } catch (error) {
        console.error('Error fetching suggestions:', error);
    }
}

function search(input) {
    searchInput.value = input;
    setTimeout(() => document.getElementById('searchForm').submit(), 300)
}

function displaySuggestions(suggestions) {
    seen = new Set()

    suggestionsContainer.innerHTML = '';

    suggestions.forEach(suggestion => {
        if (!seen.has(suggestion.Title + suggestion.Type)) {
            const div = document.createElement('div');
            div.className = 'suggestion-item';
            if (suggestion.Type == '- artist/band') {
                div.innerHTML = `
                <a class="unstyled-link" href="/artists/id=${suggestion.ID}">
                    <p>${suggestion.Title} ${suggestion.Type}</p>
                </a>
                `;
            } else {
                div.innerHTML = `<span onclick="search('${suggestion.Title}')">${suggestion.Title} ${suggestion.Type}</span>`;
            }

            suggestionsContainer.appendChild(div);
            seen.add(suggestion.Title + suggestion.Type)
        }

    });

    suggestionsContainer.style.display = 'block';
}


searchInput.addEventListener('input', function () {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
        fetchSuggestions(this.value);
    }, 300);
})

searchInput.addEventListener('keypress', function (e) {
    if (e.key === 'Enter') {
        document.getElementById('searchForm').submit();
    }
});

// ----- Filter -----

// creationDate
const minc = document.querySelector(".minc")
const maxc = document.querySelector(".maxc")
const minvc = document.querySelector(".fromCreation")
const maxvc = document.querySelector(".toCreation")
const trackc = document.querySelector(".trackCreation")

function updateTrack() {
    const minValue = parseInt(minc?.value);
    const maxValue = parseInt(maxc?.value);
    const minPercent = ((minValue - 1958) / (2015 - 1958)) * 100;
    const maxPercent = ((maxValue - 1958) / (2015 - 1958)) * 100;

    trackc.style.background = `linear-gradient(to right, lightblue ${minPercent}%, #cdcdcd ${minPercent}%, #cdcdcd ${maxPercent}%, lightblue ${maxPercent}%)`;
}

minc.addEventListener('input', () => {
    if (minc.value > maxc.value) {
        minc.value = maxc.value
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

// firstAlbum
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
