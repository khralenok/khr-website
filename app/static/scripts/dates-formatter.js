'use srict'

/**
 * Split an input string and format it in a human readable way
 */
function parseDateString(string) {
  const [day, mon, year, hm] = string.split(' ');
  const [hour, min] = hm.split(':');
  const months = { Jan:0, Feb:1, Mar:2, Apr:3, May:4, Jun:5, Jul:6, Aug:7, Sep:8, Oct:9, Nov:10, Dec:11 };

  return new Date(Date.UTC(year, months[mon], day, hour, min));
}

/**
 * Replace original string by formated one
 */
const formatTheDate = function(date, options){
        const rawDate = parseDateString(date.innerHTML)
        const formatedDate = new Intl.DateTimeFormat('en-GB', options).format(rawDate);
        date.innerHTML = " " + formatedDate
}


/**
 * Get elements from the page and call formatting funcion for each of them
 */
const registerDates = function(){
    const dates = document.querySelectorAll('.date');
    const userTimeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;

    const options = {
        day: "2-digit",
        month: "short",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        hour12: false,
        timeZone: userTimeZone, 
    }

    dates.forEach((date) => formatTheDate(date, options))
}


document.addEventListener('DOMContentLoaded', registerDates)